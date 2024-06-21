package controller

import (
	"Ayala-Crea/ResepBe/model"
	repo "Ayala-Crea/ResepBe/repository"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetCategori(c *fiber.Ctx) error {
	// Use the correct type for categories; it should be a slice of model.Categories
	var categories []model.Categories

	// Retrieve the database connection from the fiber context
	db := c.Locals("db").(*gorm.DB)

	// Fetch all categories using the repository function
	categories, err := repo.GetAllCategories(db) // Corrected function name
	if err != nil {
		// Return an error response if there was an issue fetching categories
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve categories: " + err.Error()})
	}

	// Create a response map with success information and the fetched categories
	response := fiber.Map{
		"code":    http.StatusOK,
		"success": true,
		"status":  "success",
		"data":    categories,
	}

	// Return the response as JSON
	return c.Status(http.StatusOK).JSON(response)
}

func CreateCategory(c *fiber.Ctx) error {
	var category model.Categories

	// Parse the request body into the category object
	if err := c.BodyParser(&category); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON: " + err.Error(),
		})
	}

	// Get the database connection from the Fiber context
	db := c.Locals("db").(*gorm.DB)

	// Insert the new category using the repository function
	insertedCategory, err := repo.InsertCategory(db, category)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert category: " + err.Error(),
		})
	}

	// Return the created category as a response
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"code":    http.StatusCreated,
		"success": true,
		"status":  "success",
		"message": "Category successfully saved",
		"data":    insertedCategory,
	})
}

func GetCategoriById(c *fiber.Ctx) error {
	id := c.Query("category_id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID categori tidak boleh kosong"})
	}

	db := c.Locals("db").(*gorm.DB)

	category, err := repo.GetCategoriById(db, id)
	if err != nil {
		// Jika terjadi kesalahan, mengembalikan respons error
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Data Tidak Ditemukan"})
	}

	return c.JSON(fiber.Map{"code": http.StatusOK, "success": true, "status": "success", "data": category})

}