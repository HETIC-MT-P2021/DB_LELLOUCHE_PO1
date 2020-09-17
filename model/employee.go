package model

import (
	"database/sql"
	"db_lellouche_po1/database"
	"fmt"
)

//Employee is an employee model
type Employee struct {
	Number         int                `json:"employee_number,omitempty"`
	LastName       string             `json:"last_name"`
	FirstName      string             `json:"first_name"`
	Extension      string             `json:"extension"`
	Email          string             `json:"email"`
	OfficeCode     string             `json:"office_code,omitempty"`
	ReportsTo      database.NullInt64 `json:"reports_to,omitempty"`
	JobTitle       string             `json:"job_title"`
	OfficeLocation string             `json:"office_location"`
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

//GetEmployeesWithOffice is requesting all employees and gives office additional info for each employee
func (repository *Repository) GetEmployeesWithOffice() ([]*Employee, error) {
	rows, err := repository.Conn.Query(`SELECT
		e.lastName,
		e.firstName,
		e.extension,
		e.email,
		e.jobTitle,
		o.city as officeLocation

	FROM employees e
	INNER JOIN offices o ON e.officeCode = o.officeCode`)
	if err != nil {
		return nil, fmt.Errorf("could not prepare query: %v", err)
	}

	var employees []*Employee

	var lastName, firstName, extension, email, jobTitle, officeLocation string

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&lastName, &firstName, &extension, &email, &jobTitle, &officeLocation)
		if err == sql.ErrNoRows {
			return employees, nil
		}
		if err != nil {
			return nil, fmt.Errorf("could not scan rows: %v", err)
		}

		employee := &Employee{
			LastName:       lastName,
			FirstName:      firstName,
			Extension:      extension,
			Email:          email,
			JobTitle:       jobTitle,
			OfficeLocation: officeLocation,
		}

		employees = append(employees, employee)
	}

	return employees, nil
}
