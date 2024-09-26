package tests

import (
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

func (s *GuestSessionInventoryAccessTestSuite) TestGuestInventoryAccess() {
	userJson := `{"AccountName":"Name forAccount", "APIKey":"stringthatisapikey", "Password":"stringthatispassword"}`
	req, _ := http.NewRequest("POST", "/addkey", strings.NewReader(userJson))

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = req
	s.AccountHandler.CreateGuest(c)

	cookie := w.Result().Cookies()

	fmt.Println("COOKIE", cookie[0].Value)
	assert.Equal(s.T(), "sessionID", cookie[0].Name, "Correct cookie name")

	// userID, exists := c.Get("userID")
	// assert.True(s.T(), exists, "add userid to context")
	// fmt.Println(userID)

	assert.Equal(s.T(), 200, w.Code)
}

// func (s *GuestSessionInventoryAccessTestSuite) TestNoCookieNoInventoryAccess() {}
