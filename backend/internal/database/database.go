package database

import (
	"errors"

	"log"

	"github.com/zoehay/gw2armoury/backend/internal/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/zoehay/gw2armoury/backend/internal/models"
)

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

func PostgresInit(dsn string) (*gorm.DB, error){

	// Add logic to ping db 
	// time.Sleep(30 * time.Second) 

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error initializing database connection", err)
	}

	log.Print("Run db migrate")
	err = db.AutoMigrate(&models.Item{})
    if err != nil {
        return nil, err
    }

	return db, nil

}

func CheckAndSeedDatabase(itemRepository repository.GormItemRepository) (error) {
 	_ , err := itemRepository.GetFirst()
    if errors.Is(err, gorm.ErrRecordNotFound) {
        log.Print("Seeding database")
        for _, seedItem := range seedItems {
            if _, err := itemRepository.Create(seedItem); err != nil {
                return err
            }
        }
    
    } else {
        log.Print("Database already seeded")
    }

	return nil
}
