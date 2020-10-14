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

type SimplifiedEmployee struct {
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
}

//Office is an office model
type Office struct {
	Code      int64                 `json:"office_code"`
	City      string                `json:"city"`
	Country   string                `json:"country"`
	Employees []*SimplifiedEmployee `json:"employees"`
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

func (repository *Repository) GetOfficeEmployeesInfos(ID int64) (*Office, error) {
	rows, err := repository.Conn.Query(`SELECT
		o.city,
		o.country,
		e.firstName,
		e.lastName
	FROM
		offices o
		INNER JOIN employees e ON o.officeCode = e.officeCode
	WHERE
		o.officeCode = (?);`, ID)
	if err != nil {
		return nil, fmt.Errorf("could not prepare query: %v", err)
	}

	var city, country, firstName, lastName string
	var simplifiedEmployees []*SimplifiedEmployee
	var office *Office

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&city, &country, &firstName, &lastName)
		if err == sql.ErrNoRows {
			return office, nil
		}
		if err != nil {
			return nil, fmt.Errorf("could not scan rows: %v", err)
		}

		simplifiedEmployee := &SimplifiedEmployee{
			LastName:  lastName,
			FirstName: firstName,
		}

		simplifiedEmployees = append(simplifiedEmployees, simplifiedEmployee)
	}

	office = &Office{
		Code:      ID,
		City:      city,
		Country:   country,
		Employees: simplifiedEmployees,
	}

	return office, nil
}
