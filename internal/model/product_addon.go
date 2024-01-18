package model

type ProductAddon struct {
	ID             uint64  `json:"id"`
	ProductID      uint64  `json:"product_id"`
	ProductAddonID uint64  `json:"product_addon_id"`
	Name           string  `json:"name"`
	Price          float64 `json:"price"`
}
