package tests

import (
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

	//assert multiple character names
	assert.Equal(s.T(), 200, w.Code)
}

func (s *BagItemHandlerTestSuite) TestGetByCharacterName() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/account/characters/Roman%20Meows/inventory", nil)
	req.AddCookie(s.Cookie)
	s.Router.ServeHTTP(w, req)

	//assert correct character name
	assert.Equal(s.T(), 200, w.Code)
}
