package usersessiontest

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/zoehay/gw2armoury/backend/internal/api/handlers"
	"github.com/zoehay/gw2armoury/backend/internal/api/routes"
	"github.com/zoehay/gw2armoury/backend/internal/db/repositories"
	"github.com/zoehay/gw2armoury/backend/internal/services"
)

var Router *gin.Engine
var Repository *repositories.Repository
var Service *services.Service
var AccountHandler *handlers.AccountHandler

func TestUserSession(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "User Session Suite")
}

var _ = BeforeSuite(func() {
	envPath := filepath.Join("../..", ".env")
	err := godotenv.Load(envPath)
	Expect(err).NotTo(HaveOccurred())
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	dsn := os.Getenv("TEST_DB_DSN")
	_, Repository, Service, err = routes.SetupRouter(dsn, true)
	Expect(err).NotTo(HaveOccurred())
	if err != nil {
		log.Fatal("Error setting up router", err)
	}

	AccountHandler = handlers.NewAccountHandler(&Repository.AccountRepository, &Repository.SessionRepository, Service.AccountService, Service.CharacterService)

})

var _ = AfterSuite(func() {
	db, err := Repository.AccountRepository.DB.DB()
	Expect(err).NotTo(HaveOccurred())
	db.Close()
})

var _ = AfterEach(func() {
	err := Repository.AccountRepository.DB.Exec("DROP TABLE db_accounts cascade;").Error
	Expect(err).NotTo(HaveOccurred())
	err = Repository.AccountRepository.DB.Exec("DROP TABLE db_sessions cascade;").Error
	Expect(err).NotTo(HaveOccurred())
})
