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
	gw2api "github.com/zoehay/gw2armoury/backend/internal/gw2_api"
	"github.com/zoehay/gw2armoury/backend/internal/repository"
	"github.com/zoehay/gw2armoury/backend/internal/services"
)

type CharacterServiceTestSuite struct {
	suite.Suite
	CharacterService services.CharacterService
}

func TestCharacterServiceTestSuite(t *testing.T) {
	suite.Run(t, new(CharacterServiceTestSuite))
}

func (s *CharacterServiceTestSuite) SetupSuite() {
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

	bagItemRepository := repository.NewGORMBagItemRepository(db)
	characterProvider := &gw2api.CharacterProviderMock{}
	s.CharacterService = *services.NewCharacterService(&bagItemRepository, characterProvider)
}

func (s *CharacterServiceTestSuite) TearDownSuite() {
	err := s.CharacterService.GORMBagItemRepository.DB.Exec("DROP TABLE gorm_bag_items;").Error
	assert.NoError(s.T(), err, "Failed to clear database")

	db, err := s.CharacterService.GORMBagItemRepository.DB.DB()
	if err != nil {
		s.T().Fatal(err)
	}
	db.Close()
}

func (s *CharacterServiceTestSuite) TestGetAndStoreAllCharacters() {
	err := s.CharacterService.GetAndStoreAllCharacters("accountid", "apikeystring")
	assert.NoError(s.T(), err, "Failed to get and store items")
}
