package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zoehay/gw2armoury/backend/internal/api/models"
	dbmodels "github.com/zoehay/gw2armoury/backend/internal/db/models"
	"github.com/zoehay/gw2armoury/backend/internal/db/repositories"
)

type BagItemHandler struct {
	BagItemRepository repositories.BagItemRepositoryInterface
}

func NewBagItemHandler(bagItemRepository repositories.BagItemRepositoryInterface) *BagItemHandler {
	return &BagItemHandler{
		BagItemRepository: bagItemRepository,
	}
}

func (bagItemHandler BagItemHandler) GetByCharacter(c *gin.Context) {
	characterName := c.Params.ByName("charactername")
	value, exists := c.Get("accountID")
	if !exists {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "could not find Gin Context accountID"})
		return
	}
	accountID := value.(string)

	dbItems, err := bagItemHandler.BagItemRepository.GetDetailBagItemByCharacterName(accountID, characterName)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	items := make([]models.BagItem, len(dbItems))
	for i := range dbItems {
		items[i] = dbItems[i].ToBagItem()
	}

	c.IndentedJSON(http.StatusOK, items)
}

func (bagItemHandler BagItemHandler) GetByAccount(c *gin.Context) {
	value, exists := c.Get("accountID")
	if !exists {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "could not find Gin Context accountID"})
		return
	}

	accountID := value.(string)
	dbItems, err := bagItemHandler.BagItemRepository.GetDetailBagItemByAccountID(accountID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	items := make([]models.BagItem, len(dbItems))
	for i := range dbItems {
		items[i] = dbItems[i].ToBagItem()
	}

	c.IndentedJSON(http.StatusOK, items)
}

func (bagItemHandler BagItemHandler) GetAccountInventory(c *gin.Context) {
	value, exists := c.Get("accountID")
	if !exists {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "could not find Gin Context accountID"})
		return
	}
	accountID := value.(string)

	detailBagItems, err := bagItemHandler.BagItemRepository.GetDetailBagItemByAccountID(accountID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error getting account inventory": err.Error()})
		return
	}

	accountInventory := dbmodels.DBDetailBagItemsToAccountInventory(detailBagItems, accountID)

	c.IndentedJSON(http.StatusOK, accountInventory)
}
