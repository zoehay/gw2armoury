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

type ItemServiceTestSuite struct {
	suite.Suite
	Router     *gin.Engine
	Repository *repositories.Repository
	Service    *services.Service
}

func TestItemServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ItemServiceTestSuite))
}

func (s *ItemServiceTestSuite) SetupSuite() {
	router, repository, service, err := testutils.DBRouterSetup()
	if err != nil {
		s.T().Errorf("Error setting up router: %v", err)
	}

	s.Router = router
	s.Repository = repository
	s.Service = service
}

func (s *ItemServiceTestSuite) TearDownSuite() {
	err := s.Service.ItemService.ItemRepository.DB.Exec("DROP TABLE db_items;").Error
	assert.NoError(s.T(), err, "Failed to clear database")

	db, err := s.Service.ItemService.ItemRepository.DB.DB()
	if err != nil {
		s.T().Fatal(err)
	}
	db.Close()
}

func (s *ItemServiceTestSuite) TestGetAndStoreAllItems() {
	err := s.Service.ItemService.GetAndStoreAllItems()
	assert.NoError(s.T(), err, "Failed to get and store items")
}

func (s *ItemServiceTestSuite) TestGetItemById() {
	item, err := s.Service.ItemService.ItemRepository.GetById(27952)
	assert.NoError(s.T(), err, "Failed to get item by id")
	assert.Equal(s.T(), "Axiquiotl", item.Name, "Correct item name")

}
