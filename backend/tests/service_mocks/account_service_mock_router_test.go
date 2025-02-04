package servicemocks_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/zoehay/gw2armoury/backend/internal/api/handlers"
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
	router, repository, service, err := testutils.DBRouterSetup()
	if err != nil {
		s.T().Errorf("Error setting up router: %v", err)
	}

	s.Router = router
	s.Repository = repository
	s.Service = service
	s.AccountHandler = handlers.NewAccountHandler(&repository.AccountRepository, &repository.SessionRepository, service.AccountService, service.CharacterService)
}

func (s *AccountRouterServiceTestSuite) TearDownSuite() {
	dropTables := []string{"db_accounts", "db_sessions"}
	err := testutils.TearDownDropTables(s.Repository, dropTables)
	if err != nil {
		s.T().Errorf("Error tearing down suite: %v", err)
	}
}

func (s *AccountRouterServiceTestSuite) TestGetAccount() {
	account, err := s.Service.AccountService.GetAccount("apiKey")
	testutils.PrintObject(account)
	assert.NoError(s.T(), err, "Failed to get account")
}
