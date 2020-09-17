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

//GetOrderProducts is for getting all products from a specific order
func GetOrderProducts(writer http.ResponseWriter, request *http.Request) {
	db := database.DBCon
	repository := model.Repository{Conn: db}

	muxVars := mux.Vars(request)
	ID, err := strconv.ParseInt(muxVars["id"], 10, 64)
	if err != nil {
		log.Printf("could not parse string to int: %v", err)
		helper.WriteJSONError(writer, http.StatusBadRequest, "could not parse url")
		return
	}

	products, err := repository.GetOrderItems(ID)
	if err != nil {
		log.Printf("could not get items from order with id: %d, %v", ID, err)
		helper.WriteJSONError(writer, http.StatusInternalServerError, "could not get items from order")
		return
	}
	if products == nil {
		helper.WriteJSON(writer, http.StatusNotFound, "no order with this id")
		return
	}

	helper.WriteJSON(writer, http.StatusOK, products)
}
