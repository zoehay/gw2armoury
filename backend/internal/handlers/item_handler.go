package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zoehay/gw2armoury/backend/internal/repository"
)

type ItemHandler struct {
	ItemRepository repository.GormItemRepository
}

func NewItemHandler(itemRepository repository.GormItemRepository) (*ItemHandler) {
	return &ItemHandler{
		ItemRepository: itemRepository,
	}
}

func (itemHandler ItemHandler) GetAllItems(c *gin.Context) {
	items, err := itemHandler.ItemRepository.GetAll()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, items)
}

func (itemHandler ItemHandler) GetItemByID(c *gin.Context) {
	itemId := c.Params.ByName("id")

	item, err := itemHandler.ItemRepository.GetById(itemId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, item)
}