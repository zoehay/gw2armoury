package repositories

import "gorm.io/gorm"

type RepositoryInterface interface {
	// Migrate() error
	// GetAll() (interface{}, error)
	// GetById(id int) (interface{}, error)
	// Create(entity interface{}) (interface{}, error)
}

type Repository struct {
	ItemRepository    ItemRepository
	BagItemRepository BagItemRepository
	AccountRepository AccountRepository
	SessionRepository SessionRepository
}

func NewRepository(db *gorm.DB) *Repository {
	itemRepository := NewItemRepository(db)
	bagItemRepository := NewBagItemRepository(db)
	accountRepository := NewAccountRepository(db)
	sessionRepository := NewSessionRepository(db)

	return &Repository{
		ItemRepository:    itemRepository,
		BagItemRepository: bagItemRepository,
		AccountRepository: accountRepository,
		SessionRepository: sessionRepository,
	}
}
