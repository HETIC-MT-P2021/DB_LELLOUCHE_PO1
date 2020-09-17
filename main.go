package main

import (
	"db_lellouche_po1/database"
	"db_lellouche_po1/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	log.Printf("connected to database")

	port := 8000
	newRouter := router.NewRouter()

	log.Printf("server started on port %d", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), newRouter); err != nil {
		log.Fatalf("could not initiate server: %v", err)
	}
}
