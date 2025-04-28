package servicemocks_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/zoehay/gw2armoury/backend/internal/db/repositories"
	"github.com/zoehay/gw2armoury/backend/internal/services"
	"github.com/zoehay/gw2armoury/backend/tests/testutils"
)

type BagItemAccountServiceTestSuite struct {
	suite.Suite
	Router     *gin.Engine
	Repository *repositories.Repository
	Service    *services.Service
}

func TestBagItemAccountServiceTestSuite(t *testing.T) {
	suite.Run(t, new(BagItemAccountServiceTestSuite))
}

func (s *BagItemAccountServiceTestSuite) SetupSuite() {
	router, repository, service, err := testutils.DBRouterSetup()
	if err != nil {
		s.T().Errorf("Error setting up router: %v", err)
	}

	s.Router = router
	s.Repository = repository
	s.Service = service
}

func (s *BagItemAccountServiceTestSuite) TearDownSuite() {
	dropTables := []string{"db_accounts", "db_sessions", "db_bag_items", "db_items"}
	err := testutils.TearDownDropTables(s.Repository, dropTables)
	if err != nil {
		s.T().Errorf("Error tearing down suite: %v", err)
	}
}

func (s *BagItemAccountServiceTestSuite) TestGetAndStoreAccountInventory() {
	err := s.Service.BagItemService.GetAndStoreAccountInventory("accountid", "apikeystring")
	assert.NoError(s.T(), err, "Failed to get and store items")
}

func (s *BagItemAccountServiceTestSuite) TestGetBagItemsByAccount() {
	bagItems, err := s.Service.BagItemService.BagItemRepository.GetDetailBagItemByAccountID("accountid")
	testutils.PrintObject(bagItems)
	assert.NoError(s.T(), err, "Failed to get account bag items")
}
