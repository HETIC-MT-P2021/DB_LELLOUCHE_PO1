package model

import (
	"database/sql"
	"db_lellouche_po1/database"
	"fmt"
	"time"
)

//Customer is a customer model
type Customer struct {
	Number                 int64                `json:"customer_number,omitempty"`
	Name                   string               `json:"customer_name"`
	LastName               string               `json:"contact_last_name"`
	FirstName              string               `json:"first_name"`
	Phone                  string               `json:"phone"`
	AddressLine1           string               `json:"address_line_1"`
	AddressLine2           database.NullString  `json:"address_line_2"`
	City                   string               `json:"city"`
	State                  database.NullString  `json:"state"`
	PostalCode             database.NullString  `json:"postal_code"`
	Country                string               `json:"country"`
	SalesRepEmployeeNumber database.NullInt64   `json:"sales_rep_employee_number"`
	CreditLimit            database.NullFloat64 `json:"credit_limit"`
}

//GetCustomerInfos is for requesting all infos of a specific customer
func (repository *Repository) GetCustomerInfos(ID int64) (*Customer, error) {
	row := repository.Conn.QueryRow("SELECT c.customerName, c.contactLastName, c.contactFirstName, c.phone, c.addressLine1, c.addressLine2, c.city, c.`state`, c.postalCode, c.country, c.salesRepEmployeeNumber, c.creditLimit FROM customers c WHERE c.customerNumber = (?)", ID)

	var name, lastName, firstName, phone, addressLine1, city, country string
	var state, postalCode, addressLine2 database.NullString
	var salesRepEmployeeNumber database.NullInt64
	var creditLimit database.NullFloat64

	switch err := row.Scan(&name, &lastName, &firstName, &phone, &addressLine1, &addressLine2, &city, &state, &postalCode, &country, &salesRepEmployeeNumber, &creditLimit); err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return &Customer{
			Name:                   name,
			LastName:               lastName,
			FirstName:              firstName,
			Phone:                  phone,
			AddressLine1:           addressLine1,
			AddressLine2:           addressLine2,
			City:                   city,
			State:                  state,
			PostalCode:             postalCode,
			Country:                country,
			SalesRepEmployeeNumber: salesRepEmployeeNumber,
			CreditLimit:            creditLimit,
		}, nil
	default:
		return nil, fmt.Errorf("could not scan row: %v", err)
	}
}

//GetCustomerOrders is for requesting all orders from a specific customer
func (repository *Repository) GetCustomerOrders(ID int64) ([]*Order, error) {
	rows, err := repository.Conn.Query(`SELECT
		o.orderNumber,
		o.orderDate,
		o.requiredDate,
		o.shippedDate,
		o.status,
		o.comments,
		SUM(od.priceEach * od.quantityOrdered) as priceTotal,
		SUM(quantityOrdered) as totalQuantityOrdered
	FROM
	orders o
	INNER JOIN orderdetails od ON o.orderNumber = od.orderNumber
	WHERE
	o.customerNumber = (?)
	GROUP BY
	o.orderNumber`, ID)
	if err != nil {
		return nil, fmt.Errorf("could not prepare query: %v", err)
	}

	var orders []*Order

	var number, totalQuantityOrdered int64
	var date, requiredDate time.Time
	var status string
	var priceTotal float32
	var shippedDate database.NullTime
	var comments database.NullString

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&number, &date, &requiredDate, &shippedDate, &status, &comments, &priceTotal, &totalQuantityOrdered)
		if err == sql.ErrNoRows {
			return orders, nil
		}
		if err != nil {
			return nil, fmt.Errorf("could not scan rows: %v", err)
		}

		order := &Order{
			Number:               number,
			Date:                 date,
			RequiredDate:         requiredDate,
			ShippedDate:          shippedDate,
			Status:               status,
			Comments:             comments,
			CustomerNumber:       ID,
			PriceTotal:           priceTotal,
			TotalQuantityOrdered: totalQuantityOrdered,
		}

		orders = append(orders, order)
	}

	return orders, nil
}
