package models

// Category struct represents a product category in the application.
type Category struct {
	ID   int    `json:"id"`   // Unique identifier for the category
	Name string `json:"name"` // Name of the category
}

// In-memory storage for categories
var Categories = []Category{}
