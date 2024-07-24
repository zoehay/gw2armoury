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
)

type DBTestSuite struct {
	suite.Suite
	dsn string
}

func (suite *DBTestSuite) SetupTest() {
	envPath := filepath.Join("..", ".env")
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	suite.dsn = os.Getenv("TEST_DB_DSN")
}

func (suite *DBTestSuite) TestPostgresInit() {

	_, err := database.PostgresInit(suite.dsn)
	assert.NoError(suite.T(), err, "Failed to connect to database")

}
func TestDBTestSuite(t *testing.T) {
	suite.Run(t, new(DBTestSuite))
}
