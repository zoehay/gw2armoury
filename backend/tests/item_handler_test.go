package tests

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/zoehay/gw2armoury/backend/cmd"
	"github.com/zoehay/gw2armoury/backend/internal/db/repositories"
	"github.com/zoehay/gw2armoury/backend/internal/services"
)

type ItemHandlerTestSuite struct {
	suite.Suite
	Router     *gin.Engine
	Repository *repositories.Repository
	Service    *services.Service
}

func TestItemHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(ItemHandlerTestSuite))
}

func (s *ItemHandlerTestSuite) SetupSuite() {
	envPath := filepath.Join("..", ".env")
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	dsn := os.Getenv("TEST_DB_DSN")
	router, repository, service, err := cmd.SetupTestRouter(dsn, true)
	if err != nil {
		log.Fatal("Error setting up router", err)
	}

	s.Router = router
	s.Repository = repository
	s.Service = service

	err = s.Service.ItemService.GetAndStoreAllItems()
	if err != nil {
		log.Fatal("Error getting and storing items", err)
	}

}

func (s *ItemHandlerTestSuite) TearDownSuite() {
	err := s.Service.ItemService.ItemRepository.DB.Exec("DROP TABLE db_items;").Error
	assert.NoError(s.T(), err, "Failed to clear database")

	db, err := s.Service.ItemService.ItemRepository.DB.DB()
	if err != nil {
		s.T().Fatal(err)
	}
	db.Close()
}

func (s *ItemHandlerTestSuite) TestGetItemById() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/items/27952", nil)
	s.Router.ServeHTTP(w, req)

	fmt.Println(w.Body.String())
	assert.Equal(s.T(), 200, w.Code)
}
