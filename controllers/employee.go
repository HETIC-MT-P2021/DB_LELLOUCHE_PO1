package controllers

import (
	"db_lellouche_po1/database"
	"db_lellouche_po1/helper"
	"db_lellouche_po1/model"
	"log"
	"net/http"
)

//GetEmployeesWithOffice is for getting all employees and gives office additional info for each employee
func GetEmployeesWithOffice(writer http.ResponseWriter, request *http.Request) {
	db := database.DBCon
	repository := model.Repository{Conn: db}

	employees, err := repository.GetEmployeesWithOffice()
	if err != nil {
		log.Printf("could not get all employees: %v", err)
		helper.WriteJSONError(writer, http.StatusInternalServerError, "could not get employees in DB")
		return
	}

	helper.WriteJSON(writer, http.StatusOK, employees)
}
