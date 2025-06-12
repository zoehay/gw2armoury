package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	dbmodels "github.com/zoehay/gw2armoury/backend/internal/db/models"
	"github.com/zoehay/gw2armoury/backend/internal/db/repositories"
	"github.com/zoehay/gw2armoury/backend/internal/services"
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

func (handler AccountHandler) HandlePostAccountRequest(c *gin.Context) {

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

	var requestAccount = &dbmodels.DBAccount{
		AccountID:      *gw2Account.ID,
		AccountName:    accountRequest.AccountName,
		GW2AccountName: gw2Account.Name,
		APIKey:         &accountRequest.APIKey,
		Password:       accountRequest.Password,
	}

	// determine new or returning user, return new or updated account
	account, session, err := handler.AccountService.GenerateOrUpdateAccount(requestAccount, *gw2Account)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error generating or updating account": err.Error()})
		return
	}

	c.SetCookie("sessionID", session.SessionID, 3600, "/", "localhost", false, true)

	if handler.AccountService.IsRecrawlDue(account.LastCrawl) {
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

	fmt.Println("FINISHED")
	fmt.Println(account.DBAccountToAccount())
	c.IndentedJSON(http.StatusOK, account.DBAccountToAccount())
}

func (handler AccountHandler) Delete(c *gin.Context) {

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
	_, _, err = handler.AccountService.RenewOrGenerateSession(account)
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
