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

//GetCustomerInfos is for getting all infos from a specific customer
func GetCustomerInfos(writer http.ResponseWriter, request *http.Request) {
	db := database.DBCon
	repository := model.Repository{Conn: db}

	muxVars := mux.Vars(request)
	ID, err := strconv.ParseInt(muxVars["id"], 10, 64)
	if err != nil {
		log.Printf("could not parse string to int: %v", err)
		helper.WriteJSONError(writer, http.StatusBadRequest, "could not parse url")
		return
	}

	customer, err := repository.GetCustomerInfos(ID)
	if err != nil {
		log.Printf("could not get customer with id: %d, %v", ID, err)
		helper.WriteJSONError(writer, http.StatusInternalServerError, "could not get customer")
		return
	}
	if customer == nil {
		helper.WriteJSON(writer, http.StatusNotFound, "no customer with this id")
		return
	}

	helper.WriteJSON(writer, http.StatusOK, customer)
}

//GetCustomerOrders is for getting all orders from a specific customer
func GetCustomerOrders(writer http.ResponseWriter, request *http.Request) {
	db := database.DBCon
	repository := model.Repository{Conn: db}

	muxVars := mux.Vars(request)
	ID, err := strconv.ParseInt(muxVars["id"], 10, 64)
	if err != nil {
		log.Printf("could not parse string to int: %v", err)
		helper.WriteJSONError(writer, http.StatusBadRequest, "could not parse url")
		return
	}

	orders, err := repository.GetCustomerOrders(ID)
	if err != nil {
		log.Printf("could not get orders from customer with id: %d, %v", ID, err)
		helper.WriteJSONError(writer, http.StatusInternalServerError, "could not get orders from customer")
		return
	}
	if orders == nil {
		helper.WriteJSON(writer, http.StatusNotFound, "no customer with this id")
		return
	}

	helper.WriteJSON(writer, http.StatusOK, orders)
}
