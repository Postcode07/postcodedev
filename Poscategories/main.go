package main

import (
	"Poscategories/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterCategoryRoutes(router)

	log.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", router)
}
