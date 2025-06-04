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
	gw2models "github.com/zoehay/gw2armoury/backend/internal/gw2_client/models"
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

func (handler AccountHandler) GetAccount(c *gin.Context) {

	accountID := c.MustGet("accountID").(string)
	account, err := handler.AccountRepository.GetByID(accountID)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, account.DBAccountToAccount())
}

func (handler AccountHandler) PostAccountRequest(c *gin.Context) {

	var accountRequest AccountRequest

	if err := c.BindJSON(&accountRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"request body bind json error": err.Error()})
		return
	}

	// verify GW2 account
	gw2Account, err := handler.AccountService.GetAccount(accountRequest.APIKey)
	if err != nil || gw2Account == nil || gw2Account.ID == nil || gw2Account.Name == nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"error could not get account id from gw2 api": err.Error()})
		return
	}

	// determine new or returning user, return new or updated account
	account, err := handler.GenerateOrUpdateAccount(accountRequest, *gw2Account)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error generating or updating account": err.Error()})
		return
	}

	account, session, err := handler.renewOrGenerateSession(account)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error generating or updating session": err.Error()})
		return
	}

	c.SetCookie("sessionID", session.SessionID, 3600, "/", "localhost", false, true)

	if handler.isRecrawlDue(account.LastCrawl) {
		err = handler.BagItemService.GetAndStoreAllBagItems(account.AccountID, accountRequest.APIKey)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error getting inventory after guest creation": err.Error()})
			return
		}
		err = handler.AccountRepository.UpdateLastCrawl(account.AccountID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error updating account last crawl": err.Error()})
			return
		}
	}

	c.IndentedJSON(http.StatusOK, account.DBAccountToAccount())
}

func (handler AccountHandler) GenerateOrUpdateAccount(accountRequest AccountRequest, gw2Account gw2models.GW2Account) (*dbmodels.DBAccount, error) {

	gw2AccountID := gw2Account.ID

	var account *dbmodels.DBAccount

	var newAccount = &dbmodels.DBAccount{
		AccountID:      *gw2Account.ID,
		AccountName:    accountRequest.AccountName,
		GW2AccountName: gw2Account.Name,
		APIKey:         &accountRequest.APIKey,
		Password:       accountRequest.Password,
	}

	existingAccount, err := handler.AccountRepository.GetByID(*gw2AccountID)
	if err != nil {
		// new user
		if errors.Is(err, gorm.ErrRecordNotFound) {
			account, err = handler.AccountRepository.Create(newAccount)
			if err != nil {
				return nil, fmt.Errorf("account repository create error: %s", err)
			}
		} else {
			if err != nil {
				return nil, fmt.Errorf("error accessing account db: %s", err)
			}
		}
	} else {
		// returning user
		// TODO replace password with user
		if existingAccount.Password != nil {
			// existing full account
			return nil, fmt.Errorf("error existing account for account id: %s", *gw2AccountID)
		} else {
			if newAccount.Password != nil {
				// existing guest account, accountRequest has password so upgrade to full account
				account, err = handler.AccountRepository.Update(existingAccount, newAccount)
				// TODO add password encryption
				if err != nil {
					return nil, fmt.Errorf("account repository update account error: %s", err)
				}
			} else {
				// existing guest account, no password in request so update api key
				// account, err = handler.AccountRepository.UpdateAPIKey(existingAccount.AccountID, *newAccount.APIKey)
				// if err != nil {
				// 	return nil, fmt.Errorf("account repository update apikey error: %s", err)
				// }
				account = existingAccount
			}
		}
	}

	return account, nil
}

func (handler AccountHandler) Delete(c *gin.Context) {
	fmt.Println("DELETE key")
	// use request later for User with multiple Accounts
	var deleteKeyRequest DeleteKeyRequest

	if err := c.BindJSON(&deleteKeyRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"request body bind json error": err.Error()})
		return
	}

	accountID := c.MustGet("accountID").(string)
	sessionID := c.MustGet("sessionID").(string)

	// delete api key
	err := handler.AccountRepository.DeleteAccount(accountID)
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

	// if no user (only one apikey) delete the session
	err = handler.SessionRepository.Delete(sessionID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error session items": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"API key deleted": deleteKeyRequest.APIKey})
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

	c.IndentedJSON(http.StatusOK, account.DBAccountToAccount())

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

	if account.SessionID != nil {
		session, err = handler.SessionRepository.Renew(*account.SessionID)
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

// func (handler AccountHandler) renewSession(sessionID string) (updatedSession *dbmodels.DBSession, err error) {
// 	updatedSession, err = handler.SessionRepository.Renew(sessionID)
// 	if err != nil {
// 		return nil, fmt.Errorf("renewSession error updating session: %s", err)
// 	}
// 	return updatedSession, nil
// }

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

type AccountRequest struct {
	AccountName *string
	APIKey      string
	Password    *string
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
