package handlerroutes_test

import (
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
	"github.com/zoehay/gw2armoury/backend/internal/api/models"
	"github.com/zoehay/gw2armoury/backend/internal/api/routes"
	"github.com/zoehay/gw2armoury/backend/internal/db/repositories"
	"github.com/zoehay/gw2armoury/backend/internal/services"
	"github.com/zoehay/gw2armoury/backend/tests/testutils"
)

type CreateGuestSessionTestSuite struct {
	suite.Suite
	Router     *gin.Engine
	Repository *repositories.Repository
	Service    *services.Service
}

func TestCreateGuestSessionTestSuite(t *testing.T) {
	suite.Run(t, new(CreateGuestSessionTestSuite))
}

func (s *CreateGuestSessionTestSuite) SetupSuite() {
	envPath := filepath.Join("../..", ".env")
	err := godotenv.Load(envPath)
	if err != nil {
		s.T().Errorf("Error loading .env file: %v", err)
	}

	dsn := os.Getenv("TEST_DB_DSN")
	router, repository, service, err := routes.SetupRouter(dsn, true)
	if err != nil {
		s.T().Errorf("Error setting up router: %v", err)
	}

	s.Router = router
	s.Repository = repository
	s.Service = service
}

func (s *CreateGuestSessionTestSuite) TearDownSuite() {
	err := s.Repository.AccountRepository.DB.Exec("DROP TABLE db_accounts cascade;").Error
	assert.NoError(s.T(), err, "Failed to clear database")

	err = s.Repository.AccountRepository.DB.Exec("DROP TABLE db_sessions cascade;").Error
	assert.NoError(s.T(), err, "Failed to clear database")

	err = s.Repository.AccountRepository.DB.Exec("DROP TABLE db_bag_items cascade;").Error
	assert.NoError(s.T(), err, "Failed to clear database")

	db, err := s.Repository.AccountRepository.DB.DB()
	if err != nil {
		s.T().Fatal(err)
	}
	db.Close()
}

func (s *CreateGuestSessionTestSuite) TestCreateGuestWithNewAPIKey() {
	gin.SetMode(gin.TestMode)

	userJson := `{"APIKey":"stringapikey"}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/apikeys", strings.NewReader(userJson))
	s.Router.ServeHTTP(w, req)
	assert.Equal(s.T(), 200, w.Code)

	dbAccount, err := s.Repository.AccountRepository.GetByID("gw2apiaccountidstring")
	if err != nil {
		s.T().Errorf("Error getting account from db: %v", err)
	}

	account, err := testutils.UnmarshalToType[models.Account](w)
	if err != nil {
		s.T().Errorf("Failed to unmarshal response: %v", err)
	}

	cookieSessionID := w.Result().Cookies()[0].Value

	assert.Equal(s.T(), dbAccount.SessionID, account.SessionID, "SessionID in db matches returned account")
	assert.Equal(s.T(), dbAccount.SessionID, &cookieSessionID, "SessionID in db matches returned cookie")

}
