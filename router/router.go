package router

import (
	"db_lellouche_po1/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

//Route defining a standard route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

//NewRouter is for creating a new route
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		Name:        "Get customer by ID",
		Method:      "GET",
		Pattern:     "/customers/{id}",
		HandlerFunc: controllers.GetCustomerInfos,
	},
	Route{
		Name:        "Get orders by customer ID",
		Method:      "GET",
		Pattern:     "/customers/{id}/orders/",
		HandlerFunc: controllers.GetCustomerOrders,
	},
	Route{
		Name:        "Get products by order ID",
		Method:      "GET",
		Pattern:     "/orders/{id}/products",
		HandlerFunc: controllers.GetOrderProducts,
	},
}
