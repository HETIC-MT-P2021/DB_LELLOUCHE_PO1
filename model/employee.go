package model

import (
	"db_lellouche_po1/database"
)

//Employee is an employee model
type Employee struct {
	Number     int                `json:"employee_number"`
	LastName   string             `json:"last_name"`
	FirstName  string             `json:"first_name"`
	Extension  string             `json:"extension"`
	Email      string             `json:"email"`
	OfficeCode string             `json:"office_code"`
	ReportsTo  database.NullInt64 `json:"reports_to"`
	JobTitle   string             `json:"job_title"`
}

//Office is an office model
type Office struct {
	Code         string              `json:"office_code"`
	City         string              `json:"city"`
	Phone        string              `json:"phone"`
	AddressLine1 string              `json:"address_line_1"`
	AddressLine2 database.NullString `json:"address_line_2"`
	State        database.NullString `json:"state"`
	Country      string              `json:"country"`
	PostalCode   string              `json:"postal_code"`
	Territory    string              `json:"territory"`
}
