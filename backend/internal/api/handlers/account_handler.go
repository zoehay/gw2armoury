package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zoehay/gw2armoury/backend/internal/database/repository"
	repositorymodels "github.com/zoehay/gw2armoury/backend/internal/database/repository_models"
	"github.com/zoehay/gw2armoury/backend/internal/services"
)

type AccountHandler struct {
	AccountRepository repository.AccountRepositoryInterface
	SessionRepository repository.SessionRepositoryInterface
	AccountService    services.AccountServiceInterface
}

func NewAccountHandler(accountRepository repository.AccountRepositoryInterface, sessionRepository repository.SessionRepositoryInterface, accountService services.AccountServiceInterface) *AccountHandler {
	return &AccountHandler{
		AccountRepository: accountRepository,
		SessionRepository: sessionRepository,
		AccountService:    accountService,
	}
}

func (handler AccountHandler) Create(c *gin.Context) {
	var accountCreate AccountCreate

	if err := c.BindJSON(&accountCreate); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	apiAccountID, err := handler.AccountService.GetAccountID(accountCreate.APIKey)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error could not get account id from gw2 api": err.Error()})
		return
	}

	var stringAccountID string
	if apiAccountID != nil {
		stringAccountID = *apiAccountID
	}

	var newAccount = &repositorymodels.DBAccount{
		AccountID:   stringAccountID,
		AccountName: &accountCreate.AccountName,
	}

	account, err := handler.AccountRepository.Create(newAccount)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = handler.startSession(c, account)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

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
	err = handler.startSession(c, account)
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

func (handler AccountHandler) startSession(c *gin.Context, account *repositorymodels.DBAccount) error {
	// Create a session
	newSessionID := generateSessionID()
	var newSession = &repositorymodels.DBSession{
		SessionID: newSessionID,
		Expires:   time.Now(),
	}

	session, err := handler.SessionRepository.Create(newSession)
	if err != nil {
		return err
	}

	// Add session to account
	_, err = handler.AccountRepository.UpdateSession(account.AccountID, session.SessionID)
	if err != nil {
		return err
	}

	c.SetCookie("sessionID", session.SessionID, 3600, "/", "localhost", false, true)
	return nil
}

func generateSessionID() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

type AccountLogin struct {
	AccountName string
	Password    string
}

type AccountCreate struct {
	AccountName string
	APIKey      string
}
