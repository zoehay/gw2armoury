package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/zoehay/gw2armoury/backend/internal/handlers"
	"github.com/zoehay/gw2armoury/backend/internal/models"
	"github.com/zoehay/gw2armoury/backend/internal/repository"
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

    itemRepository := repository.NewGormItemRepository(db)
    itemHandler := handlers.NewItemHandler(itemRepository)

    router := gin.Default()
    router.GET("/items", itemHandler.GetAllItems) 
    router.GET("/items/:id", itemHandler.GetItemByID)

    // router.Run("127.0.0.1:8000")
    router.Run(":8000")
}

func postgresInit() (*gorm.DB, error) {

    // Add logic to ping db 
    time.Sleep(30 * time.Second) 

    dsn := os.Getenv("DB_DSN")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    
    if err != nil {
        return nil, err
    }

    // itemRepository := &repository.GormItemRepository{DB: db}

    // log.Print("Run db migrate")
	// err = db.AutoMigrate(&models.Item{})
    // if err != nil {
    //     return nil, err
    // }

    // seed db 

    // _ , err = itemRepository.GetFirst()
    // if errors.Is(err, gorm.ErrRecordNotFound) {
    //     log.Print("Seeding db")
    //     for _, seedItem := range seedItems {
    //         if _, err := itemRepository.Create(seedItem); err != nil {
    //             return db, err
    //         }
    //     }
    
    // } else {
    //     log.Print("db already seeded")
    // }

    return db, nil
}