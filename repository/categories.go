package repository

import (
	"Ayala-Crea/ResepBe/model"

	"gorm.io/gorm"
)

func GetAllCategories(db *gorm.DB) ([]model.Categories, error) { // Corrected function name to be consistent
	var categories []model.Categories // Use the plural form to indicate multiple categories

	// Use GORM to find all entries in the categories table
	if err := db.Find(&categories).Error; err != nil {
		return nil, err
	}

	// Return the fetched categories and a nil error (indicating success)
	return categories, nil
}

func InsertCategory(db *gorm.DB, category model.Categories) (*model.Categories, error) {
	if err := db.Create(&category).Error; err != nil {
		return nil, err // Return the error if creation fails
	}

	return &category, nil // Return the created category on success
}

func GetCategoriById(db *gorm.DB, id string) (model.Categories, error) {
	var category model.Categories
	if err := db.First(&category, "category_id = ?", id).Error; err != nil {
		return category, err
	}
	return category, nil
}
