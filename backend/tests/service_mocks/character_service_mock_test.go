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
	"github.com/zoehay/gw2armoury/backend/internal/gw2_client/providers"
	"github.com/zoehay/gw2armoury/backend/internal/services"
	"github.com/zoehay/gw2armoury/backend/tests/testutils"
)

type CharacterServiceTestSuite struct {
	suite.Suite
	CharacterService services.CharacterService
}

func TestCharacterServiceTestSuite(t *testing.T) {
	suite.Run(t, new(CharacterServiceTestSuite))
}

func (s *CharacterServiceTestSuite) SetupSuite() {
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

	bagItemRepository := repositories.NewBagItemRepository(db)
	characterProvider := &providers.CharacterProviderMock{}
	s.CharacterService = *services.NewCharacterService(&bagItemRepository, characterProvider)
}

func (s *CharacterServiceTestSuite) TearDownSuite() {
	err := s.CharacterService.BagItemRepository.DB.Exec("DROP TABLE db_bag_items;").Error
	assert.NoError(s.T(), err, "Failed to clear database")

	db, err := s.CharacterService.BagItemRepository.DB.DB()
	if err != nil {
		s.T().Fatal(err)
	}
	db.Close()
}

func (s *CharacterServiceTestSuite) TestGetAndStoreAllCharacters() {
	err := s.CharacterService.GetAndStoreAllCharacters("accountid", "apikeystring")
	assert.NoError(s.T(), err, "Failed to get and store items")
}

func (s *CharacterServiceTestSuite) TestGetBagItemsByCharacterName() {
	bagItems, err := s.CharacterService.BagItemRepository.GetByCharacterName("Roman Meows")
	fmt.Println(testutils.PrintObject(bagItems[0]))
	assert.NoError(s.T(), err, "Failed to get item by id")

}
