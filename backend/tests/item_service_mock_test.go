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
	itemService *servicemocks.ItemServiceMock
}

func (suite *ItemServiceTestSuite) SetupTest() {
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
	suite.itemService = servicemocks.NewItemServiceMock(&itemRepository)
}

func (suite *ItemServiceTestSuite) TestGetAndStoreAllItems() {
	err := suite.itemService.GetAndStoreAllItems()
	assert.NoError(suite.T(), err, "Failed to get and store items")
}

func TestItemServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ItemServiceTestSuite))
}
