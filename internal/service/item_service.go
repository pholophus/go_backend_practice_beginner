package service

import(
	"github.com/pholophus/go_backend_practice_beginner/internal/models"
	"github.com/pholophus/go_backend_practice_beginner/internal/repository"
)

type ItemService interface{
	GetItems() []models.Item
	GetItem(id int) (models.Item, bool)
	CreateItem(item models.Item) models.Item
	UpdateItem(item models.Item) (models.Item, bool)
	DeleteItem(id int) bool
}

type itemService struct {
	repo repository.ItemRepository
}

func NewItemService(repo repository.ItemRepository) ItemService {
	return &itemService{
		repo: repo,
	}
}

func (s *itemService) GetItems() []models.Item {
	return s.repo.GetAll()
}

func (s *itemService) GetItem(id int) (models.Item, bool) {
	return s.repo.GetByID(id)
}

func (s *itemService) CreateItem(item models.Item) models.Item {
	return s.repo.Create(item)
}

func (s *itemService) UpdateItem(item models.Item) (models.Item, bool) {
	return s.repo.Update(item)
}

func (s *itemService) DeleteItem(id int) bool {
	return s.repo.Delete(id)
}

