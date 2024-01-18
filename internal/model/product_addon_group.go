package model

type ProductAddonGroup struct {
	ID             uint64 `json:"id"`
	ProductID      uint64 `json:"product_id"`
	Name           string `json:"name"`
	MaximumQuantiy uint   `json:"maximum_quantity"`
}
