package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zoehay/gw2armoury/backend/internal/repository"
)

type BagItemHandler struct {
	BagItemRepository repository.GormBagItemRepository
}

func NewBagItemHandler(bagItemRepository repository.GormBagItemRepository) *BagItemHandler {
	return &BagItemHandler{
		BagItemRepository: bagItemRepository,
	}
}

func (BagItemHandler BagItemHandler) GetBagItemsByCharacter(c *gin.Context) {
	characterName := c.Params.ByName("charactername")
	fmt.Println("name", characterName)
	items, err := BagItemHandler.BagItemRepository.GetByCharacterName(characterName)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, items)
}
