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
)

var db *gorm.DB
var err error

type Item struct {
    ID uint `gorm:"primaryKey"`
    ChatLink string
    Name string
    Icon string
    Description string
    Type string
    Rarity string
    Level uint
    VendorValue uint
    DefaultSkin uint
    Flags []string `gorm:"type:text"`
    GameTypes []string `gorm:"type:text"`
    Restrictions []string `gorm:"type:text"`
    UpgradesInto []string `gorm:"type:text"`
    UpgradesFrom []string `gorm:"type:text"`
    Details string;
  }

var seedItems = []*Item{
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
    // router.GET("/items/:id", getItemByID)

    // router.Run("127.0.0.1:8000")
    router.Run("0.0.0.0:8000")
}

func getItems(c *gin.Context) {
    var allItems []Item
    result := db.Find(&allItems)

    if result.Error != nil {
        c.AbortWithStatus(404)
        fmt.Println(result.Error)
    } 
    c.IndentedJSON(http.StatusOK, allItems)
}

func postgresInit() (*gorm.DB, error) {

    time.Sleep(30 * time.Second)

    dsn := os.Getenv("DB_DSN")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    
    if err != nil {
        return nil, err
    }

    log.Print("Run db migrate")
	err = db.AutoMigrate(&Item{})
    if err != nil {
        return nil, err
    }

    // seed db 
    err = db.First(&Item{}).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        log.Print("Seeding db")
        result := db.Create(&seedItems)
        log.Print(result.Error, result.RowsAffected)
    } else {
        log.Print("db already seeded")
    }

    return db, nil
}