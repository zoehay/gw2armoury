package servicemocks_test

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	database "github.com/zoehay/gw2armoury/backend/internal/db"
	"github.com/zoehay/gw2armoury/backend/internal/db/repositories"
	providers "github.com/zoehay/gw2armoury/backend/internal/gw2_client/providers"
	"github.com/zoehay/gw2armoury/backend/internal/services"
	"github.com/zoehay/gw2armoury/backend/tests/testutils"
)

type ItemServiceTestSuite struct {
	suite.Suite
	ItemService services.ItemService
}

func TestItemServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ItemServiceTestSuite))
}

func (s *ItemServiceTestSuite) SetupSuite() {
	envPath := filepath.Join("../..", ".env")
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	dsn := os.Getenv("TEST_DB_DSN")
	db, err := database.PostgresInit(dsn)
	if err != nil {
		log.Fatal("Error connecting to postgres", err)
	}

	itemRepository := repositories.NewItemRepository(db)
	itemProvider := &providers.ItemProviderMock{}
	s.ItemService = *services.NewItemService(&itemRepository, itemProvider)
}

func (s *ItemServiceTestSuite) TearDownSuite() {
	err := s.ItemService.ItemRepository.DB.Exec("DROP TABLE db_items;").Error
	assert.NoError(s.T(), err, "Failed to clear database")

	db, err := s.ItemService.ItemRepository.DB.DB()
	if err != nil {
		s.T().Fatal(err)
	}
	db.Close()
}

func (s *ItemServiceTestSuite) TestGetAndStoreAllItems() {
	err := s.ItemService.GetAndStoreAllItems()
	assert.NoError(s.T(), err, "Failed to get and store items")
}

func (s *ItemServiceTestSuite) TestGetItemById() {
	item, err := s.ItemService.ItemRepository.GetById(27952)
	fmt.Println(testutils.PrintObject(item))
	assert.NoError(s.T(), err, "Failed to get item by id")
}
