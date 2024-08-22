package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func (BagItemHandler BagItemHandler) GetByCharacter(c *gin.Context) {
	characterName := c.Params.ByName("charactername")

	items, err := BagItemHandler.BagItemRepository.GetIconBagItemByCharacterName(characterName)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, items)
}

// func (BagItemHandler BagItemHandler) GetByAccount(c *gin.Context) {
// 	accountId := c.Params.ByName("accountid")
// 	items, err := BagItemHandler.BagItemRepository.GetDetailsByAccountID(accountID)
// 	if err != nil {
// 		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.IndentedJSON(http.StatusOK, items)
// }
