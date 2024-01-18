package service

import (
	"automatic-guacamole/internal/model"
	"automatic-guacamole/internal/repository"

	"gorm.io/gorm"
)

type ProductAddonGroupService interface {
	GetAll() ([]model.ProductAddonGroup, error)
	GetByID(id uint64) (*model.ProductAddonGroup, error)
	Create(product *model.ProductAddonGroup) error
	Update(product *model.ProductAddonGroup) error
	Delete(id uint64) error
}

type productAddonGroupService struct {
	productAddonGroupRepo repository.ProductAddonGroupRepository
	db                    *gorm.DB
}

func NewProductAddonGroupService(db *gorm.DB, repo repository.ProductAddonGroupRepository) ProductAddonGroupService {
	return &productAddonGroupService{productAddonGroupRepo: repo, db: db}
}

func (s *productAddonGroupService) GetAll() ([]model.ProductAddonGroup, error) {
	return s.productAddonGroupRepo.GetAll(s.db)
}

func (s *productAddonGroupService) GetByID(id uint64) (*model.ProductAddonGroup, error) {
	return s.productAddonGroupRepo.GetByID(s.db, id)
}

func (s *productAddonGroupService) Create(product *model.ProductAddonGroup) error {
	return s.productAddonGroupRepo.Create(s.db, product)
}

func (s *productAddonGroupService) Update(product *model.ProductAddonGroup) error {
	return s.productAddonGroupRepo.Update(s.db, product)
}

func (s *productAddonGroupService) Delete(id uint64) error {
	return s.productAddonGroupRepo.Delete(s.db, id)
}
