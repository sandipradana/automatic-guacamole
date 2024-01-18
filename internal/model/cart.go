package model

type Cart struct {
	Items []CartItem `json:"items"`
}

type CartItem struct {
	ProductID uint64          `json:"product_id"`
	Addons    []CartItemAddon `json:"addons"`
}

type CartItemAddon struct {
	AddonID uint64 `json:"addon_id"`
}
