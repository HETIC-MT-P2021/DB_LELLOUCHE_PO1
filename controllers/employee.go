package controllers

import (
	"db_lellouche_po1/database"
	"db_lellouche_po1/helper"
	"db_lellouche_po1/model"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
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

func GetOfficeEmployeesInfos(writer http.ResponseWriter, request *http.Request) {
	db := database.DBCon
	repository := model.Repository{Conn: db}

	muxVars := mux.Vars(request)
	ID, err := strconv.ParseInt(muxVars["id"], 10, 64)
	if err != nil {
		log.Printf("could not parse string to int: %v", err)
		helper.WriteJSONError(writer, http.StatusBadRequest, "could not parse url")
		return
	}

	office, err := repository.GetOfficeEmployeesInfos(ID)
	if err != nil {
		log.Printf("could not get office: %v", err)
		helper.WriteJSONError(writer, http.StatusInternalServerError, "could not get office in DB")
		return
	}
	if office == nil {
		helper.WriteJSON(writer, http.StatusNotFound, "no customer with this id")
		return
	}

	helper.WriteJSON(writer, http.StatusOK, office)
}
