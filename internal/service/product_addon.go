package service

import (
	"automatic-guacamole/internal/model"
	"automatic-guacamole/internal/repository"

	"gorm.io/gorm"
)

type ProductAddonService interface {
	GetAll() ([]model.ProductAddon, error)
	GetByID(id uint64) (*model.ProductAddon, error)
	Create(product *model.ProductAddon) error
	Update(product *model.ProductAddon) error
	Delete(id uint64) error
}

type productAddonService struct {
	productAddonRepo repository.ProductAddonRepository
	db               *gorm.DB
}

func NewProductAddonService(db *gorm.DB, repo repository.ProductAddonRepository) ProductAddonService {
	return &productAddonService{productAddonRepo: repo, db: db}
}

func (s *productAddonService) GetAll() ([]model.ProductAddon, error) {
	return s.productAddonRepo.GetAll(s.db)
}

func (s *productAddonService) GetByID(id uint64) (*model.ProductAddon, error) {
	return s.productAddonRepo.GetByID(s.db, id)
}

func (s *productAddonService) Create(product *model.ProductAddon) error {
	return s.productAddonRepo.Create(s.db, product)
}

func (s *productAddonService) Update(product *model.ProductAddon) error {
	return s.productAddonRepo.Update(s.db, product)
}

func (s *productAddonService) Delete(id uint64) error {
	return s.productAddonRepo.Delete(s.db, id)
}
