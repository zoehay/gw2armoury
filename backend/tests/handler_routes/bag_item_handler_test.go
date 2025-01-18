package tests

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/zoehay/gw2armoury/backend/internal/api/models"
	"github.com/zoehay/gw2armoury/backend/internal/api/routes"
	"github.com/zoehay/gw2armoury/backend/internal/db/repositories"
	"github.com/zoehay/gw2armoury/backend/internal/services"
)

type BagItemHandlerTestSuite struct {
	suite.Suite
	Router     *gin.Engine
	Repository *repositories.Repository
	Service    *services.Service
	Cookie     *http.Cookie
}

func TestBagItemHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(BagItemHandlerTestSuite))
}

func (s *BagItemHandlerTestSuite) SetupSuite() {
	envPath := filepath.Join("../..", ".env")
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	dsn := os.Getenv("TEST_DB_DSN")
	router, repository, service, err := routes.SetupRouter(dsn, true)
	if err != nil {
		log.Fatal("Error setting up router", err)
	}

	s.Router = router
	s.Repository = repository
	s.Service = service

	s.T().Log("Create account with POST /apikeys")
	userJson := `{"AccountName":"Name forAccount", "APIKey":"stringthatisapikey", "Password":"stringthatispassword"}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/apikeys", strings.NewReader(userJson))
	s.Router.ServeHTTP(w, req)
	assert.Equal(s.T(), http.StatusOK, w.Code)
	cookie := w.Result().Cookies()

	s.Cookie = cookie[0]
}

func (s *BagItemHandlerTestSuite) TearDownSuite() {
	err := s.Repository.AccountRepository.DB.Exec("DROP TABLE db_accounts cascade;").Error
	assert.NoError(s.T(), err, "Failed to clear database")

	err = s.Repository.AccountRepository.DB.Exec("DROP TABLE db_sessions cascade;").Error
	assert.NoError(s.T(), err, "Failed to clear database")

	err = s.Repository.AccountRepository.DB.Exec("DROP TABLE db_bag_items;").Error
	assert.NoError(s.T(), err, "Failed to clear database")

	db, err := s.Repository.AccountRepository.DB.DB()
	if err != nil {
		s.T().Fatal(err)
	}
	db.Close()
}

func (s *BagItemHandlerTestSuite) TestGetByAccount() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/account/characters/inventory", nil)
	req.AddCookie(s.Cookie)
	s.Router.ServeHTTP(w, req)

	response, err := UnmarshalBagItems(w.Body.Bytes())
	if err != nil {
		s.T().Fatalf("Failed to unmarshal response: %v", err)
	}
	assert.Equal(s.T(), 200, w.Code)

	bagItemsResponseOK := BagItemsResponseOK(response)
	assert.Equal(s.T(), true, bagItemsResponseOK)

	allSameCharacterName := BagItemsAllSameCharacterName(response)
	assert.Equal(s.T(), false, allSameCharacterName)
}

func (s *BagItemHandlerTestSuite) TestGetByCharacterName() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/account/characters/Roman%20Meows/inventory", nil)
	req.AddCookie(s.Cookie)
	s.Router.ServeHTTP(w, req)

	response, err := UnmarshalBagItems(w.Body.Bytes())
	if err != nil {
		s.T().Fatalf("Failed to unmarshal response: %v", err)
	}

	bagItemsResponseOK := BagItemsResponseOK(response)
	assert.Equal(s.T(), true, bagItemsResponseOK)

	allSameCharacterName := BagItemsAllSameCharacterName(response)
	assert.Equal(s.T(), true, allSameCharacterName)
	assert.Equal(s.T(), 200, w.Code)
}

func UnmarshalBagItems(bodyBytes []byte) ([]models.BagItem, error) {
	var response []models.BagItem
	err := json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func PrintObject(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func BagItemsResponseOK(bagItems []models.BagItem) bool {
	if len(bagItems) == 0 {
		return false
	} else {
		if len(bagItems[0].CharacterName) != 0 {
			return true
		} else {
			return false
		}
	}
}

func BagItemsAllSameCharacterName(bagItems []models.BagItem) bool {
	characterName := bagItems[0].CharacterName
	for _, bagItem := range bagItems {
		if bagItem.CharacterName == characterName {
			continue
		} else {
			return false
		}
	}
	return true
}
