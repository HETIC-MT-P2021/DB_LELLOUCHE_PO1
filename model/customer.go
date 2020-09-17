package model

import "database/sql"

type Customer struct {
	Number         int64           `json:"customer_number"`
	Name           string          `json:"customer_name"`
	LastName        string          `json:"contact_last_name"`
	Phone                  string          `json:"phone"`
	AddressLine1           string          `json:"adress_line_1"`
	AddressLine2           string          `json:"address_line_2"`
	City                   string          `json:"city"`
	State                  string          `json:"state"`
	PostalCode             string          `json:"postal_code"`
	Country                string          `json:"country"`
	SalesRepEmployeeNumber int64           `json:"sales_rep_employee_number"`
	CreditLimit            sql.NullFloat64 `json:"credit_limit"`
}

func (repository *Repository) GetCustomerInfos(customer *Customer) error {

	return nil
}
