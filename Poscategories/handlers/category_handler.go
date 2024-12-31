package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"Poscategories/models"
)

func AddCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	category.ID = len(models.Categories) + 1
	models.Categories = append(models.Categories, category)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Categories)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedCategory models.Category
	if err := json.NewDecoder(r.Body).Decode(&updatedCategory); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	for i, category := range models.Categories {
		if category.ID == id {
			models.Categories[i].Name = updatedCategory.Name
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(models.Categories[i])
			return
		}
	}
	http.Error(w, "Category not found", http.StatusNotFound)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i, category := range models.Categories {
		if category.ID == id {
			models.Categories = append(models.Categories[:i], models.Categories[i+1:]...)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Category deleted"))
			return
		}
	}
	http.Error(w, "Category not found", http.StatusNotFound)
}
