package service

import (
	"automatic-guacamole/internal/model"
	"automatic-guacamole/internal/repository"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type CartService interface {
	Calculate(cart *model.Cart) (float64, error)
}

type cartService struct {
	productRepo           repository.ProductRepository
	productAddonGroupRepo repository.ProductAddonGroupRepository
	productAddonRepo      repository.ProductAddonRepository
	db                    *gorm.DB
}

func NewCartService(db *gorm.DB, productRepo repository.ProductRepository, productAddonGroupRepo repository.ProductAddonGroupRepository, productAddonRepo repository.ProductAddonRepository) CartService {
	return &cartService{productRepo: productRepo, db: db, productAddonGroupRepo: productAddonGroupRepo, productAddonRepo: productAddonRepo}
}

func (s *cartService) Calculate(cart *model.Cart) (float64, error) {

	var totalPrice float64 = 0

	var productTotalAddon map[string]uint = map[string]uint{}

	for i := range cart.Items {
		product, _ := s.productRepo.GetByID(s.db, cart.Items[i].ProductID)
		if product != nil {
			totalPrice = totalPrice + product.Price
			for j := range cart.Items[i].Addons {
				productAddon, _ := s.productAddonRepo.GetByID(s.db, cart.Items[i].Addons[j].AddonID)
				if productAddon != nil {
					productAddonGroup, _ := s.productAddonGroupRepo.GetByID(s.db, cart.Items[i].Addons[j].AddonID)
					if productAddonGroup != nil {
						productAddonGroupKey := fmt.Sprintf("%d-%d", product.ID, productAddonGroup.ID)
						_, ok := productTotalAddon[productAddonGroupKey]
						if ok {
							productTotalAddon[productAddonGroupKey] = productTotalAddon[productAddonGroupKey] + 1

							if productTotalAddon[productAddonGroupKey] > productAddonGroup.MaximumQuantiy {
								return 0, errors.New("addon max")
							}
						} else {
							productTotalAddon[productAddonGroupKey] = 1
						}
					}

					totalPrice = totalPrice + productAddon.Price
				}
			}
		}
	}

	return totalPrice, nil
}
