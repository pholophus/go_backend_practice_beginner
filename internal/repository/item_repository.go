package repository

import "github.com/pholophus/go_backend_practice_beginner/internal/models"

type ItemRepository interface {
	GetAll() []models.Item
	GetByID(id int) (models.Item, bool)
	Create(item models.Item) models.Item
	Update(item models.Item) (models.Item, bool)
	Delete(id int) bool
}

type inMemoryItemRepository struct {
	items []models.Item
	nextID int
}

func NewInMemoryRepository() ItemRepository{
	return &inMemoryItemRepository{
		item: make([]models.Item, 0),
		nextID: 1,
	}
}

func (r *inMemoryItemRepository) GetAll() []models.Item {
	return r.items
}

func (r *inMemoryItemRepository) GetByID(id int) (models.Item, bool) {
	for _, item := range r.items {
		if item.ID == id {
			return item, true
		}
	}
	return models.Item{}, false
}

func (r *inMemoryItemRepository) Create(item models.Item) models.Item {
	item.ID = r.nextID
	r.nextID++
	r.items = append(r.items, item)
	return item
}

func (r *inMemoryItemRepository) Update(updated models.Item) (models.Item, bool) {
	for i, item := range r.items {
		if item.ID == updated.ID {
			r.items[i] = updated
			return updated, true
		}
		return models.Item{}, false
	}
}

func (r *inMemoryItemRepository) Delete(id int) bool {
	for i, item := range r.items {
		if item.ID == updated.ID {
			r.items = append(r.items[:i], r.items[i+1:]...)
			return true
		}
		return false
	}
}