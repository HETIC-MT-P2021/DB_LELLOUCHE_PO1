package model

//Product is a product model
type Product struct {
	Code            string  `json:"product_code"`
	Name            string  `json:"product_name"`
	Line            string  `json:"product_line"`
	Scale           string  `json:"product_scale"`
	Vendor          string  `json:"product_vendor"`
	Description     string  `json:"product_description"`
	QuantityInStock int64   `json:"quantity_in_stock"`
	BuyPrice        float32 `json:"buy_price"`
	MSRP            float32 `json:"msrp"`
}
