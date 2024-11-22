package tests

import (
	"encoding/json"
	"fmt"
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
	"github.com/zoehay/gw2armoury/backend/internal/api/handlers"
	"github.com/zoehay/gw2armoury/backend/internal/api/routes"
	dbmodels "github.com/zoehay/gw2armoury/backend/internal/db/models"
	"github.com/zoehay/gw2armoury/backend/internal/db/repositories"
	"github.com/zoehay/gw2armoury/backend/internal/services"
)

type GuestSessionInventoryAccessTestSuite struct {
	suite.Suite
	Router         *gin.Engine
	Repository     *repositories.Repository
	Service        *services.Service
	AccountHandler *handlers.AccountHandler
}

func TestGuestSessionInventoryAccessSuite(t *testing.T) {
	suite.Run(t, new(GuestSessionInventoryAccessTestSuite))
}

func (s *GuestSessionInventoryAccessTestSuite) SetupSuite() {
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
	s.AccountHandler = handlers.NewAccountHandler(&repository.AccountRepository, &repository.SessionRepository, service.AccountService, service.CharacterService)

	s.Service.ItemService.GetAndStoreAllItems()

}

func (s *GuestSessionInventoryAccessTestSuite) TearDownSuite() {
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

func (s *GuestSessionInventoryAccessTestSuite) TestNoCookieNoInventoryAccess() {
	s.T().Log("GET /account with no cookie")
	w0 := httptest.NewRecorder()
	req0, _ := http.NewRequest("GET", "/account/characters/Roman%20Meows/inventory", nil)
	s.Router.ServeHTTP(w0, req0)
	assert.Equal(s.T(), http.StatusForbidden, w0.Code)
}

func (s *GuestSessionInventoryAccessTestSuite) TestGuestInventoryAccess() {
	s.T().Log("Create account with POST /apikeys")
	userJson := `{"AccountName":"Name forAccount", "APIKey":"stringthatisapikey", "Password":"stringthatispassword"}`
	w1 := httptest.NewRecorder()
	req1, _ := http.NewRequest("POST", "/apikeys", strings.NewReader(userJson))
	s.Router.ServeHTTP(w1, req1)
	assert.Equal(s.T(), http.StatusOK, w1.Code)
	cookie := w1.Result().Cookies()

	s.T().Log("GET /account with cookie")
	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("GET", "/account/characters/Roman%20Meows/inventory", nil)
	req2.AddCookie(cookie[0])
	s.Router.ServeHTTP(w2, req2)
	assert.Equal(s.T(), http.StatusOK, w2.Code)

	var response []dbmodels.DBIconBagItem
	err := json.Unmarshal(w2.Body.Bytes(), &response)
	if err != nil {
		s.T().Fatalf("Failed to unmarshal response: %v", err)
	}
	fmt.Println(PrintObject(response))
}

func PrintObject(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
