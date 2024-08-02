package tests

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/zoehay/gw2armoury/backend/internal/database"
	"github.com/zoehay/gw2armoury/backend/internal/repository"
	servicemocks "github.com/zoehay/gw2armoury/backend/tests/service_mocks"
)

type ItemServiceTestSuite struct {
	suite.Suite
	ItemService *servicemocks.ItemServiceMock
}

func TestItemServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ItemServiceTestSuite))
}

func (s *ItemServiceTestSuite) SetupSuite() {
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

	itemRepository := repository.NewGORMItemRepository(db)
	s.ItemService = servicemocks.NewItemServiceMock(&itemRepository)
}

func (s *ItemServiceTestSuite) TearDownSuite() {
	err := s.ItemService.GormItemRepository.DB.Exec("DROP TABLE gorm_items;").Error
	assert.NoError(s.T(), err, "Failed to clear database")

	db, err := s.ItemService.GormItemRepository.DB.DB()
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
	item, err := s.ItemService.GormItemRepository.GetById(27952)
	log.Println(item)
	assert.NoError(s.T(), err, "Failed to get item by id")
}
