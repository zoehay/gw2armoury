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
	BagItemRepository repositories.BagItemRepositoryInterface
	AccountService    services.AccountServiceInterface
	BagItemService    services.BagItemServiceInterface
}

func NewAccountHandler(accountRepository repositories.AccountRepositoryInterface, sessionRepository repositories.SessionRepositoryInterface, bagItemRepostiory repositories.BagItemRepositoryInterface, accountService services.AccountServiceInterface, bagItemService services.BagItemServiceInterface) *AccountHandler {
	return &AccountHandler{
		AccountRepository: accountRepository,
		SessionRepository: sessionRepository,
		BagItemRepository: bagItemRepostiory,
		AccountService:    accountService,
		BagItemService:    bagItemService,
	}
}

func (handler AccountHandler) PostAPIKey(c *gin.Context) {

	var apiKeyRequest APIKeyRequest

	if err := c.BindJSON(&apiKeyRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"request body bind json error": err.Error()})
		return
	}

	// verify GW2 account
	gw2Account, err := handler.AccountService.GetAccount(apiKeyRequest.APIKey)
	if err != nil || gw2Account == nil || gw2Account.ID == nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"error could not get account id from gw2 api": err.Error()})
		return
	}

	var account *dbmodels.DBAccount
	var session *dbmodels.DBSession
	gw2AccountID := gw2Account.ID

	existingAccount, err := handler.AccountRepository.GetByID(*gw2AccountID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// new user
		account, session, err = handler.generateNewGuestAccount(*gw2AccountID, apiKeyRequest.APIKey)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error creating new guest account": err.Error()})
			return
		}
	} else if err != nil {
		// database error
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error accessing account database": err.Error()})
		return
	} else {
		// existing user
		account = existingAccount
		err := handler.AccountRepository.UpdateAPIKey(*account.APIKey, apiKeyRequest.APIKey)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error updating api key": err.Error()})
		}
		account, session, err = handler.renewOrGenerateSession(account)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"session error for existing guest": err.Error()})
		}
	}

	c.SetCookie("sessionID", session.SessionID, 3600, "/", "localhost", false, true)

	if handler.isRecrawlDue(account.LastCrawl) {
		err = handler.BagItemService.GetAndStoreAllBagItems(*gw2AccountID, apiKeyRequest.APIKey)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error getting inventory after guest creation": err.Error()})
			return
		}
		err = handler.AccountRepository.UpdateLastCrawl(*gw2AccountID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error updating account last crawl": err.Error()})
			return
		}
	}

	c.IndentedJSON(http.StatusOK, account.ToAccount())
}

func (handler AccountHandler) DeleteAPIKey(c *gin.Context) {
	var deleteKeyRequest DeleteKeyRequest

	if err := c.BindJSON(&deleteKeyRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"request body bind json error": err.Error()})
		return
	}

	accountID := c.MustGet("accountID").(string)

	// delete api key from account
	err := handler.AccountRepository.DeleteAPIKey(accountID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error deleting api key": err.Error()})
		return
	}

	// delete associated bag items
	err = handler.BagItemRepository.DeleteByAccountID(accountID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error deleting bag items": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"API key deleted": deleteKeyRequest.APIKey})
}

func (handler AccountHandler) Create(c *gin.Context) {
	var createRequest CreateRequest

	if err := c.BindJSON(&createRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"request body bind json error": err.Error()})
		return
	}

	// verify GW2 account
	gw2Account, err := handler.AccountService.GetAccount(createRequest.APIKey)
	if err != nil || gw2Account == nil || gw2Account.ID == nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"error could not get account id from gw2 api": err.Error()})
		return
	}
	// TODO password encryption

	gw2AccountID := gw2Account.ID

	var account *dbmodels.DBAccount
	var session *dbmodels.DBSession
	var newAccount = &dbmodels.DBAccount{
		AccountID:   *gw2AccountID,
		AccountName: &createRequest.AccountName,
		APIKey:      &createRequest.APIKey,
		Password:    &createRequest.Password,
		// Session:     newSession,
	}

	existingAccount, err := handler.AccountRepository.GetByID(*gw2AccountID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			account, err = handler.AccountRepository.Create(newAccount)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"account repository create error": err.Error()})
				return
			}
		} else {
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"error accessing account db": err.Error()})
				return
			}
		}
	} else if existingAccount != nil {
		if existingAccount.Password != nil {
			// existing account
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error existing account for account id": err.Error()})
			return
		} else {
			// existing guest account, upgrade to full account
			account, err = handler.AccountRepository.Update(existingAccount, newAccount)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"account repository updates error": err.Error()})
				return
			}
		}
	}

	updatedAccount, session, err := handler.renewOrGenerateSession(account)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"session error for existing guest": err.Error()})
	}

	if handler.isRecrawlDue(account.LastCrawl) {
		err = handler.BagItemService.GetAndStoreAllBagItems(*gw2AccountID, *account.APIKey)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error getting inventory after guest creation": err.Error()})
			return
		}
		err = handler.AccountRepository.UpdateLastCrawl(*gw2AccountID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error updating account last crawl": err.Error()})
			return
		}
	}

	c.SetCookie("sessionID", session.SessionID, 3600, "/", "localhost", false, true)
	c.IndentedJSON(http.StatusOK, updatedAccount.ToAccount())

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

func (handler AccountHandler) renewOrGenerateSession(account *dbmodels.DBAccount) (*dbmodels.DBAccount, *dbmodels.DBSession, error) {
	var session *dbmodels.DBSession
	var err error

	if account.Session != nil {
		session, err = handler.renewSession(account.Session)
		if err != nil {
			return nil, nil, fmt.Errorf("error renewing session for existing account: %w", err)
		}
	} else {
		account, session, err = handler.generateNewSession(account)
		if err != nil {
			return nil, nil, fmt.Errorf("error generating new session for existing account: %w", err)
		}
	}

	return account, session, nil
}

func (handler AccountHandler) isRecrawlDue(lastCrawl *time.Time) bool {
	minHoursSinceCrawl := float64(1)
	var elapsed float64

	if lastCrawl != nil {
		t := time.Now()
		elapsed = t.Sub(*lastCrawl).Hours()
	}

	return (elapsed >= minHoursSinceCrawl || lastCrawl == nil)
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
		Expires:   time.Now().Add(3600 * time.Second),
	}

	newSession, err = handler.SessionRepository.Create(session)
	if err != nil {
		return nil, nil, err
	}

	updatedAccount, err = handler.AccountRepository.UpdateSession(account.AccountID, newSession)
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

type APIKeyRequest struct {
	APIKey string
}

type DeleteKeyRequest struct {
	APIKey string
}
