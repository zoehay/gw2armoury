package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/zoehay/gw2armoury/backend/internal/models"
)

var db *gorm.DB
var err error

var seedItems = []*models.Item{
    {
        ID: 28445, 
        Name: "Strong Soft Wood Longbow of Fire", 
        Icon: "https://render.guildwars2.com/file/C6110F52DF5AFE0F00A56F9E143E9732176DDDE9/65015.png", 
        Description: "",
        Type: "Weapon",},
    {
        ID: 12452, 
        Name: "Omnomberry Bar", 
        Type: "Consumable",
        Level: 80,
        Rarity: "Fine",
        Icon: "https://render.guildwars2.com/file/6BD5B65FBC6ED450219EC86DD570E59F4DA3791F/433643.png", 
    },
}

func main() {

    err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

    db, err = postgresInit()
    if err != nil {
		log.Fatal("Error initializing database connection", err)
	}

    router := gin.Default()
    router.GET("/items", getItems) 
    router.GET("/items/:id", getItemByID)

    // router.Run("127.0.0.1:8000")
    router.Run(":8000")
}

func getItems(c *gin.Context) {
    var allItems []models.Item

    err = db.Find(&allItems).Error
    if err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
        fmt.Println(err)
        return
    } 
    c.IndentedJSON(http.StatusOK, allItems)
}

func getItemByID(c *gin.Context) {
    var item models.Item
    itemID := c.Params.ByName("id")

    err = db.First(&item, itemID).Error
    if err!= nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
        fmt.Println(err)
        return
    } 
    c.IndentedJSON(http.StatusOK, item)
}

func postgresInit() (*gorm.DB, error) {

    // Add logic to ping db 
    time.Sleep(30 * time.Second) 

    dsn := os.Getenv("DB_DSN")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    
    if err != nil {
        return nil, err
    }

    log.Print("Run db migrate")
	err = db.AutoMigrate(&models.Item{})
    if err != nil {
        return nil, err
    }

    // seed db 
    err = db.First(&models.Item{}).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        log.Print("Seeding db")
        result := db.Create(&seedItems)
        log.Print(result.Error, result.RowsAffected)
    } else {
        log.Print("db already seeded")
    }

    return db, nil
}