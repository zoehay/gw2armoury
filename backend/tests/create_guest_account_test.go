package tests

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/zoehay/gw2armoury/backend/internal/api/routes"
	"github.com/zoehay/gw2armoury/backend/internal/db/repositories"
	"github.com/zoehay/gw2armoury/backend/internal/services"
)

type CreateGuestAccountTestSuite struct {
	suite.Suite
	Router     *gin.Engine
	Repository *repositories.Repository
	Service    *services.Service
}

func TestCreateGuestAccountTestSuite(t *testing.T) {
	suite.Run(t, new(CreateGuestAccountTestSuite))
}

func (s *CreateGuestAccountTestSuite) SetupSuite() {
	envPath := filepath.Join("..", ".env")
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	dsn := os.Getenv("TEST_DB_DSN")
	router, repository, service, err := routes.SetupRouter(dsn, true)
	if err != nil {
		log.Fatal("Error setting up router", err)
	}

	s.Router = router
	s.Repository = repository
	s.Service = service
}

func (s *CreateGuestAccountTestSuite) TearDownSuite() {
	err := s.Repository.AccountRepository.DB.Exec("DROP TABLE db_accounts cascade;").Error
	assert.NoError(s.T(), err, "Failed to clear database")

	err = s.Repository.AccountRepository.DB.Exec("DROP TABLE db_sessions cascade;").Error
	assert.NoError(s.T(), err, "Failed to clear database")

	db, err := s.Repository.AccountRepository.DB.DB()
	if err != nil {
		s.T().Fatal(err)
	}
	db.Close()
}

func (s *CreateGuestAccountTestSuite) TestCreateGuestWithNewAPIKey() {
	gin.SetMode(gin.TestMode)

	userJson := `{"APIKey":"stringthatisapikey"}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/apikeys", strings.NewReader(userJson))
	s.Router.ServeHTTP(w, req)

	fmt.Println(w.Body.String())
	assert.Equal(s.T(), 200, w.Code)
}
