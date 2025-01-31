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
	CharacterService  services.CharacterServiceInterface
}

func NewAccountHandler(accountRepository repositories.AccountRepositoryInterface, sessionRepository repositories.SessionRepositoryInterface, accountService services.AccountServiceInterface, characterService services.CharacterServiceInterface) *AccountHandler {
	return &AccountHandler{
		AccountRepository: accountRepository,
		SessionRepository: sessionRepository,
		AccountService:    accountService,
		CharacterService:  characterService,
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
	var session *dbmodels.DBSession
	existingAccount, err := handler.AccountRepository.GetByID(*gw2AccountID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		account, session, err = handler.generateNewGuestAccount(*gw2AccountID, createRequest.APIKey)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error creating new guest account": err.Error()})
			return
		}
	} else if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error accessing account database": err.Error()})
		return
	} else {
		account = existingAccount
		if account.Session != nil {
			session, err = handler.renewSession(account.Session)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"error renewing session for existing account": err.Error()})
				return
			}
		} else {
			account, session, err = handler.generateNewSession(account)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"error generating new session for existing account": err.Error()})
				return
			}
		}
	}

	c.SetCookie("sessionID", session.SessionID, 3600, "/", "localhost", false, true)

	err = handler.CharacterService.GetAndStoreAllCharacters(*gw2AccountID, createRequest.APIKey)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error getting characters after guest creation": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, account.ToAccount())
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

	c.IndentedJSON(http.StatusOK, account.ToAccount())

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
	_, _, err = handler.generateNewSession(account)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, account.ToAccount())

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

func (handler AccountHandler) generateNewGuestAccount(accountID string, apiKey string) (updatedAccount *dbmodels.DBAccount, newSession *dbmodels.DBSession, err error) {
	var newAccount = &dbmodels.DBAccount{
		AccountID: accountID,
		APIKey:    &apiKey,
	}

	account, err := handler.AccountRepository.Create(newAccount)
	if err != nil {
		return nil, nil, err
	}

	if err != nil {
		return nil, nil, fmt.Errorf("account repository create error: %s", err)
	}

	updatedAccount, newSession, err = handler.generateNewSession(account)
	if err != nil {
		return nil, nil, fmt.Errorf("error generating session: %s", err)
	}

	return updatedAccount, newSession, nil
}

func (handler AccountHandler) generateNewSession(account *dbmodels.DBAccount) (updatedAccount *dbmodels.DBAccount, newSession *dbmodels.DBSession, err error) {
	newSessionID, err := handler.generateSessionID()
	if err != nil {
		return nil, nil, err
	}

	var session = &dbmodels.DBSession{
		SessionID: newSessionID,
		Expires:   time.Now().Add(120 * time.Second),
	}

	newSession, err = handler.SessionRepository.Create(session)
	if err != nil {
		return nil, nil, err
	}

	updatedAccount, err = handler.AccountRepository.UpdateSession(account.AccountID, session)
	if err != nil {
		return nil, nil, err
	}

	return updatedAccount, newSession, nil
}

func (handler AccountHandler) renewSession(session *dbmodels.DBSession) (updatedSession *dbmodels.DBSession, err error) {
	updatedSession, err = handler.SessionRepository.Update(session.SessionID)
	if err != nil {
		return nil, fmt.Errorf("renewSession error updating session: %s", err)
	}
	return updatedSession, nil
}

func (handler AccountHandler) generateSessionID() (sessionID string, err error) {
	b := make([]byte, 32)
	_, err = rand.Read(b)
	if err != nil {
		return "", err
	}
	sessionID = base64.RawURLEncoding.EncodeToString(b)
	return sessionID, nil
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
