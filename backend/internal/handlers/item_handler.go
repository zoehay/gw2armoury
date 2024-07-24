package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zoehay/gw2armoury/backend/internal/repository"
)

type ItemHandler struct {
	ItemRepository repository.ItemRepository
}

func NewItemHandler(itemRepository repository.ItemRepository) *ItemHandler {
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
	stringId := c.Params.ByName("id")
	itemId, err := strconv.Atoi(stringId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	item, err := itemHandler.ItemRepository.GetById(itemId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, item)
}
