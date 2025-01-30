package servicemocks_test

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/zoehay/gw2armoury/backend/internal/api/handlers"
	"github.com/zoehay/gw2armoury/backend/internal/api/routes"
	"github.com/zoehay/gw2armoury/backend/internal/db/repositories"
	"github.com/zoehay/gw2armoury/backend/internal/services"
	"github.com/zoehay/gw2armoury/backend/tests/testutils"
)

type AccountRouterServiceTestSuite struct {
	suite.Suite
	Router         *gin.Engine
	Repository     *repositories.Repository
	Service        *services.Service
	AccountHandler *handlers.AccountHandler
}

func TestAccountRouterServiceTestSuite(t *testing.T) {
	suite.Run(t, new(AccountRouterServiceTestSuite))
}

func (s *AccountRouterServiceTestSuite) SetupSuite() {
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

func (s *AccountRouterServiceTestSuite) TearDownSuite() {
	err := s.Repository.AccountRepository.DB.Exec("DROP TABLE db_accounts;").Error
	assert.NoError(s.T(), err, "Failed to clear database")

	db, err := s.Repository.AccountRepository.DB.DB()
	if err != nil {
		s.T().Fatal(err)
	}
	db.Close()
}

func (s *AccountRouterServiceTestSuite) TestGetAccount() {
	account, err := s.Service.AccountService.GetAccount("apiKey")
	fmt.Println(testutils.PrintObject(account))
	assert.NoError(s.T(), err, "Failed to get account")
}
