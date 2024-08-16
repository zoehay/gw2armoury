package tests

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/zoehay/gw2armoury/backend/internal/database"
	gw2api "github.com/zoehay/gw2armoury/backend/internal/gw2_api"
	"github.com/zoehay/gw2armoury/backend/internal/repository"
	"github.com/zoehay/gw2armoury/backend/internal/services"
)

type AccountServiceTestSuite struct {
	suite.Suite
	AccountService services.AccountService
}

func TestAccountServiceTestSuite(t *testing.T) {
	suite.Run(t, new(AccountServiceTestSuite))
}

func (s *AccountServiceTestSuite) SetupSuite() {
	envPath := filepath.Join("..", ".env")
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	dsn := os.Getenv("TEST_DB_DSN")
	db, err := database.PostgresInit(dsn)
	if err != nil {
		log.Fatal("Error connecting to postgres", err)
	}

	accountRepository := repository.NewGORMAccountRepository(db)
	accountProvider := &gw2api.AccountProviderMock{}
	s.AccountService = *services.NewAccountService(&accountRepository, accountProvider)
}

// func (s *AccountServiceTestSuite) TearDownSuite() {
// 	err := s.AccountService.GORMAccountRepository.DB.Exec("DROP TABLE accounts;").Error
// 	assert.NoError(s.T(), err, "Failed to clear database")

// 	db, err := s.AccountService.GORMAccountRepository.DB.DB()
// 	if err != nil {
// 		s.T().Fatal(err)
// 	}
// 	db.Close()
// }

func (s *AccountServiceTestSuite) TestGetAccount() {
	account, err := s.AccountService.GetAccountID("apiKey")
	fmt.Println(PrintObject(account))
	assert.NoError(s.T(), err, "Failed to get account")
}
