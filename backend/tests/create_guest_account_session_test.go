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
	"github.com/zoehay/gw2armoury/backend/cmd"
	"github.com/zoehay/gw2armoury/backend/internal/db/repositories"
	"github.com/zoehay/gw2armoury/backend/internal/services"
)

type CreateGuestAccountSessionTestSuite struct {
	suite.Suite
	Router     *gin.Engine
	Repository *repositories.Repository
	Service    *services.Service
}

func TestCreateGuestAccountSessionSuite(t *testing.T) {
	suite.Run(t, new(CreateGuestAccountSessionTestSuite))
}

func (s *CreateGuestAccountSessionTestSuite) SetupSuite() {
	envPath := filepath.Join("..", ".env")
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	dsn := os.Getenv("TEST_DB_DSN")
	router, repository, service, err := cmd.SetupTestRouter(dsn, true)
	if err != nil {
		log.Fatal("Error setting up router", err)
	}

	s.Router = router
	s.Repository = repository
	s.Service = service
}

// func (s *StartSessionTestSuite) TearDownSuite() {
// 	err := s.ItemService.ItemRepository.DB.Exec("DROP TABLE db_accounts;").Error
// 	assert.NoError(s.T(), err, "Failed to clear database")

// 	db, err := s.ItemService.ItemRepository.DB.DB()
// 	if err != nil {
// 		s.T().Fatal(err)
// 	}
// 	db.Close()
// }

func (s *CreateGuestAccountSessionTestSuite) TestAddAPIKey() {
	// w := httptest.NewRecorder()
	// c, _ := gin.CreateTestContext(w)
	// req := &http.Request{
	// 	URL:    &url.URL{},
	// 	Header: make(http.Header),
	// }

	// q := req.URL.Query()
	// q.Add("id", "27952")
	// req.URL.RawQuery = q.Encode()

	// c.Request = req
	gin.SetMode(gin.TestMode)

	userJson := `{"AccountName":"Name forAccount", "APIKey":"stringthatisapikey", "Password":"stringthatispassword"}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/addkey", strings.NewReader(userJson))
	s.Router.ServeHTTP(w, req)

	fmt.Println(w.Body.String())
	assert.Equal(s.T(), 200, w.Code)
}
