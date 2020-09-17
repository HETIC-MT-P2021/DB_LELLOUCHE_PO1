package model

import (
	"database/sql"
	"db_lellouche_po1/database"
	"fmt"
	"time"
)

//Order is an order model
type Order struct {
	Number               int64               `json:"order_number"`
	Date                 time.Time           `json:"order_date"`
	RequiredDate         time.Time           `json:"required_date"`
	ShippedDate          database.NullTime   `json:"shipped_date"`
	Status               string              `json:"status"`
	Comments             database.NullString `json:"comments"`
	CustomerNumber       int64               `json:"customer_number"`
	PriceTotal           float32             `json:"price_total"`
	TotalQuantityOrdered int64               `json:"total_quantity_ordered"`
}

/*type OrderDetails struct {
	Number          int                  `json:"order_number"`
	ProductCode     string               `json:"product_code"`
	QuantityOrdered int                  `json:"quantity_ordered"`
	PriceEach       database.NullFloat64 `json:"price_each"`
	LineNumber      int                  `json:"order_line_number"`
}*/

//GetOrderItems is for requesting all items from a specific order
func (repository *Repository) GetOrderItems(ID int64) ([]*Product, error) {
	rows, err := repository.Conn.Query(`SELECT
		p.productCode, 
		p.productName,
		p.productLine,
		p.productScale,
		p.productVendor,
		p.productDescription,
		p.quantityInStock,
		p.buyPrice,
		p.MSRP
	FROM orderdetails od
	INNER JOIN products p ON od.productCode = p.productCode
	WHERE od.orderNumber = (?)`, ID)
	if err != nil {
		return nil, fmt.Errorf("could not prepare query: %v", err)
	}

	var products []*Product

	var code, name, line, scale, vendor, description string
	var quantityInStock int64
	var buyPrice, msrp float32

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&code, &name, &line, &scale, &vendor, &description, &quantityInStock, &buyPrice, &msrp)
		if err == sql.ErrNoRows {
			return products, nil
		}
		if err != nil {
			return nil, fmt.Errorf("could not scan rows: %v", err)
		}

		product := &Product{
			Code:            code,
			Name:            name,
			Line:            line,
			Scale:           scale,
			Vendor:          vendor,
			Description:     description,
			QuantityInStock: quantityInStock,
			BuyPrice:        buyPrice,
			MSRP:            msrp,
		}

		products = append(products, product)
	}

	return products, nil
}
