package routes

import (
	"Poscategories/handlers"

	"github.com/gorilla/mux"
)

func RegisterCategoryRoutes(router *mux.Router) {
	router.HandleFunc("/categories", handlers.AddCategory).Methods("POST")
	router.HandleFunc("/categories", handlers.GetAllCategories).Methods("GET")
	router.HandleFunc("/categories/{id}", handlers.UpdateCategory).Methods("PUT")
	router.HandleFunc("/categories/{id}", handlers.DeleteCategory).Methods("DELETE")
}
