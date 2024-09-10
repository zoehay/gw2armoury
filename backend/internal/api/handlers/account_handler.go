package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	dbmodels "github.com/zoehay/gw2armoury/backend/internal/db/models"
	"github.com/zoehay/gw2armoury/backend/internal/db/repositories"
	"github.com/zoehay/gw2armoury/backend/internal/services"
	"gorm.io/gorm"
)

type AccountHandler struct {
	AccountRepository repositories.AccountRepositoryInterface
	SessionRepository repositories.SessionRepositoryInterface
	AccountService    services.AccountServiceInterface
}

func NewAccountHandler(accountRepository repositories.AccountRepositoryInterface, sessionRepository repositories.SessionRepositoryInterface, accountService services.AccountServiceInterface) *AccountHandler {
	return &AccountHandler{
		AccountRepository: accountRepository,
		SessionRepository: sessionRepository,
		AccountService:    accountService,
	}
}

func (handler AccountHandler) CreateGuest(c *gin.Context) {

	var createRequest CreateGuestRequest

	if err := c.BindJSON(&createRequest); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"request body bind json error": err.Error()})
		return
	}

	gw2AccountID, err := handler.AccountService.GetAccountID(createRequest.APIKey)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error could not get account id from gw2 api": err.Error()})
		return
	}

	var account *dbmodels.DBAccount
	existingAccount, err := handler.AccountRepository.GetByID(*gw2AccountID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		account, err = handler.generateNewGuestAccount(*gw2AccountID, createRequest.APIKey)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error creating new guest account": err.Error()})
			return
		}
	} else if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error accessing account database": err.Error()})
		return
	} else {
		account = existingAccount
		_, err = handler.renewSession(account.Session)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error renewing session for existing account": err.Error()})
			return
		}
	}

	c.SetCookie("sessionID", account.Session.SessionID, 3600, "/", "localhost", false, true)
	c.IndentedJSON(http.StatusOK, account)
}

func (handler AccountHandler) Create(c *gin.Context) {
	var accountCreate CreateRequest

	if err := c.BindJSON(&accountCreate); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"request body bind json error": err.Error()})
		return
	}

	gw2AccountID, err := handler.AccountService.GetAccountID(accountCreate.APIKey)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error could not get account id from gw2 api": err.Error()})
		return
	}

	account, err := handler.AccountRepository.GetByID(*gw2AccountID)
	if account != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error existing account for account id": err.Error()})
		return
	}

	var stringAccountID string
	if gw2AccountID != nil {
		stringAccountID = *gw2AccountID
	}

	//password encryption

	var newAccount = &dbmodels.DBAccount{
		AccountID:   stringAccountID,
		AccountName: &accountCreate.AccountName,
		APIKey:      &accountCreate.APIKey,
		// Session:     newSession,
	}

	account, err = handler.AccountRepository.Create(newAccount)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"account repository create error": err.Error()})
		return
	}

	// err = handler.startSession(c, account)
	// if err != nil {
	// 	c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	c.IndentedJSON(http.StatusOK, account)

}

func (handler AccountHandler) Login(c *gin.Context) {
	var accountLogin AccountLogin

	if err := c.BindJSON(&accountLogin); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Find account
	account, err := handler.AccountRepository.GetByName(accountLogin.AccountName)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Add password verification
	_, err = handler.generateNewSession(account)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, account)

	// refresh account info in db

}

func (handler AccountHandler) Logout(c *gin.Context) {
	// find session
	// delete session
	sessionID, err := c.Cookie("sessionID")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = handler.SessionRepository.Delete(sessionID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// find account
	// delete session from account

	// delete cookie
	c.SetCookie("sessionID", "", -1, "/", "localhost", false, true)
}

func (handler AccountHandler) generateNewGuestAccount(accountID string, apiKey string) (account *dbmodels.DBAccount, err error) {
	var newAccount = &dbmodels.DBAccount{
		AccountID: accountID,
		APIKey:    &apiKey,
	}

	account, err = handler.AccountRepository.Create(newAccount)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, fmt.Errorf("account repository create error: %s", err)
	}

	_, err = handler.generateNewSession(account)
	if err != nil {
		return nil, fmt.Errorf("error generating session: %s", err)
	}

	return account, nil
}

func (handler AccountHandler) generateNewSession(account *dbmodels.DBAccount) (session *dbmodels.DBSession, err error) {
	newSessionID := handler.generateSessionID()
	var newSession = &dbmodels.DBSession{
		SessionID: newSessionID,
		Expires:   time.Now().Add(120 * time.Second),
	}

	session, err = handler.SessionRepository.Create(newSession)
	if err != nil {
		return nil, err
	}

	account, err = handler.AccountRepository.UpdateSession(account.AccountID, session.SessionID)
	if err != nil {
		return nil, err
	}

	return account.Session, nil
}

func (handler AccountHandler) renewSession(session *dbmodels.DBSession) (updatedSession *dbmodels.DBSession, err error) {
	updatedSession, err = handler.SessionRepository.Update(session.SessionID)
	if err != nil {
		return nil, fmt.Errorf("renewSession error updating session: %s", err)
	}
	return updatedSession, nil
}

func (handler AccountHandler) generateSessionID() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

type AccountLogin struct {
	AccountName string
	Password    string
}

type CreateRequest struct {
	AccountName string
	APIKey      string
	Password    string
}

type CreateGuestRequest struct {
	APIKey string
}
