package repository

import (
	"automatic-guacamole/internal/model"

	"gorm.io/gorm"
)

var ProductAddonTable string = "productaddons"

type ProductAddonRepository interface {
	GetAll(db *gorm.DB) ([]model.ProductAddon, error)
	GetByID(db *gorm.DB, id uint64) (*model.ProductAddon, error)
	Create(db *gorm.DB, productAddon *model.ProductAddon) error
	Update(db *gorm.DB, productAddon *model.ProductAddon) error
	Delete(db *gorm.DB, id uint64) error
}

type productAddonRepository struct {
}

func NewProductAddonRepository() ProductAddonRepository {
	return &productAddonRepository{}
}

func (r *productAddonRepository) GetAll(db *gorm.DB) ([]model.ProductAddon, error) {
	var productAddons []model.ProductAddon
	if err := db.Table(ProductAddonTable).Find(&productAddons).Error; err != nil {
		return nil, err
	}
	return productAddons, nil
}

func (r *productAddonRepository) GetByID(db *gorm.DB, id uint64) (*model.ProductAddon, error) {
	var productAddon model.ProductAddon
	if err := db.Table(ProductAddonTable).First(&productAddon, id).Error; err != nil {
		return nil, err
	}
	return &productAddon, nil
}

func (r *productAddonRepository) Create(db *gorm.DB, productAddon *model.ProductAddon) error {
	if err := db.Table(ProductAddonTable).Create(productAddon).Error; err != nil {
		return err
	}
	return nil
}

func (r *productAddonRepository) Update(db *gorm.DB, productAddon *model.ProductAddon) error {
	if err := db.Table(ProductAddonTable).Save(productAddon).Error; err != nil {
		return err
	}
	return nil
}

func (r *productAddonRepository) Delete(db *gorm.DB, id uint64) error {
	if err := db.Table(ProductAddonTable).Delete(&model.ProductAddon{}, id).Error; err != nil {
		return err
	}
	return nil
}
