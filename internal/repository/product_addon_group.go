package repository

import (
	"automatic-guacamole/internal/model"

	"gorm.io/gorm"
)

var ProductAddonGroupTable string = "productaddongroups"

type ProductAddonGroupRepository interface {
	GetAll(db *gorm.DB) ([]model.ProductAddonGroup, error)
	GetByID(db *gorm.DB, id uint64) (*model.ProductAddonGroup, error)
	Create(db *gorm.DB, productAddonGroup *model.ProductAddonGroup) error
	Update(db *gorm.DB, productAddonGroup *model.ProductAddonGroup) error
	Delete(db *gorm.DB, id uint64) error
}

type productAddonGroupRepository struct {
}

func NewProductAddonGroupRepository() ProductAddonGroupRepository {
	return &productAddonGroupRepository{}
}

func (r *productAddonGroupRepository) GetAll(db *gorm.DB) ([]model.ProductAddonGroup, error) {
	var productAddonGroups []model.ProductAddonGroup
	if err := db.Table(ProductAddonGroupTable).Find(&productAddonGroups).Error; err != nil {
		return nil, err
	}
	return productAddonGroups, nil
}

func (r *productAddonGroupRepository) GetByID(db *gorm.DB, id uint64) (*model.ProductAddonGroup, error) {
	var productAddonGroup model.ProductAddonGroup
	if err := db.Table(ProductAddonGroupTable).First(&productAddonGroup, id).Error; err != nil {
		return nil, err
	}
	return &productAddonGroup, nil
}

func (r *productAddonGroupRepository) Create(db *gorm.DB, productAddonGroup *model.ProductAddonGroup) error {
	if err := db.Table(ProductAddonGroupTable).Create(productAddonGroup).Error; err != nil {
		return err
	}
	return nil
}

func (r *productAddonGroupRepository) Update(db *gorm.DB, productAddonGroup *model.ProductAddonGroup) error {
	if err := db.Table(ProductAddonGroupTable).Save(productAddonGroup).Error; err != nil {
		return err
	}
	return nil
}

func (r *productAddonGroupRepository) Delete(db *gorm.DB, id uint64) error {
	if err := db.Table(ProductAddonGroupTable).Delete(&model.ProductAddonGroup{}, id).Error; err != nil {
		return err
	}
	return nil
}
