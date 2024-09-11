package tests

import (
	"encoding/json"
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
	"github.com/zoehay/gw2armoury/backend/internal/api/handlers"
	"github.com/zoehay/gw2armoury/backend/internal/db/repositories"
	"github.com/zoehay/gw2armoury/backend/internal/services"
)

type CreateGuestAccountSessionTestSuite struct {
	suite.Suite
	Router         *gin.Engine
	Repository     *repositories.Repository
	Service        *services.Service
	AccountHandler *handlers.AccountHandler
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
	s.AccountHandler = handlers.NewAccountHandler(&repository.AccountRepository, &repository.SessionRepository, service.AccountService)

}

func (s *CreateGuestAccountSessionTestSuite) TearDownSuite() {
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

func (s *CreateGuestAccountSessionTestSuite) TestCreateGuestWithNewAPIKey() {
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
	userJson := `{"AccountName":"Name forAccount", "APIKey":"stringthatisapikey", "Password":"stringthatispassword"}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/addkey", strings.NewReader(userJson))

	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)

	c.Request = req

	s.AccountHandler.CreateGuest(c)

	// userID, exists := c.Get("userID")
	// assert.True(s.T(), exists, "add userid to context")

	cookie := w.Result().Cookies()
	fmt.Println("COOKIE", cookie[0])
	fmt.Println("COOKIE", cookie[0].Value)
	assert.Equal(s.T(), "sessionID", cookie[0].Name)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		s.T().Fatalf("Failed to unmarshal response: %v", err)
	}

	assert.Equal(s.T(), response["SessionID"].(string), cookie[0].Value)
	assert.Equal(s.T(), 200, w.Code)
}
