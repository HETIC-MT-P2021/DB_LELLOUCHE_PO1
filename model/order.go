package model

import (
	"database/sql"
	"time"
)

type Order struct {
	Number    int       `json:"order_number"`
	Date      time.Time `json:"order_date"`
	RequiredDate   time.Time `json:"required_date"`
	ShippedDate    time.Time `json:"shipped_date"`
	Status         string    `json:"status"`
	Comments       []string  `json:"comments"`
	CustomerNumber int       `json:"customer_number"`
}

type OrderDetails struct {
	Number     int             `json:"order_number"`
	ProductCode     string          `json:"product_code"`
	QuantityOrdered int             `json:"quantity_ordered"`
	PriceEach       sql.NullFloat64 `json:"price_each"`
	LineNumber int             `json:"order_line_number"`
}
