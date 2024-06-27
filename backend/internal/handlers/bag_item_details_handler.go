package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zoehay/gw2armoury/backend/internal/repository"
)

type BagItemDetailsHandler struct {
	BagItemRepository repository.GORMBagItemRepository
}

func NewBagItemDetailsHandler(bagItemRepository repository.GORMBagItemRepository) *BagItemDetailsHandler {
	return &BagItemDetailsHandler{
		BagItemRepository: bagItemRepository,
	}
}

func (BagItemDetailsHandler BagItemDetailsHandler) GetByCharacter(c *gin.Context) {
	characterName := c.Params.ByName("charactername")

	items, err := BagItemDetailsHandler.BagItemRepository.GetDetailsByCharacterName(characterName)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, items)
}

// func (BagItemDetailsHandler BagItemDetailsHandler) GetByAccount(c *gin.Context) {
// 	accountId := c.Params.ByName("accountid")
// 	items, err := BagItemDetailsHandler.BagItemRepository.GetDetailsByAccountID(accountID)
// 	if err != nil {
// 		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.IndentedJSON(http.StatusOK, items)
// }
