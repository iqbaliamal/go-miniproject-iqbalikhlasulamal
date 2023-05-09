package controllers

import (
	"go-miniproject-iqbalikhlasulamal/database"
	"go-miniproject-iqbalikhlasulamal/models/entity"
	"go-miniproject-iqbalikhlasulamal/models/request"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// struct

func CategoryGetAll(c *fiber.Ctx) error {
	category := []entity.Category{}

	database.DB.Find(&category)

	return c.JSON(category)
}

func CategoryGetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	category := entity.Category{}

	database.DB.Find(&category, id)
	if category.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "not found",
		})
	}

	return c.JSON(category)
}

func CategoryCreate(c *fiber.Ctx) error {
	categoryRequest := new(request.CategoryRequest)

	if err := c.BodyParser(categoryRequest); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(categoryRequest)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	var category entity.Category

	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// validate unique name
	var isNameExist entity.Category
	errCheckName := database.DB.First(&isNameExist, "name = ?", categoryRequest.Name).Error
	if errCheckName == nil {
		return c.Status(402).JSON(fiber.Map{
			"message": "category name already used.",
		})
	}

	database.DB.Create(&category)

	return c.JSON(category)
}

func CategoryUpdate(c *fiber.Ctx) error {
	categoryRequest := new(request.CategoryRequest)

	if err := c.BodyParser(categoryRequest); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(categoryRequest)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	id := c.Params("id")

	var category entity.Category
	var isNameExist entity.Category

	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errCheckName := database.DB.First(&isNameExist, "name = ?", categoryRequest.Name).Error
	if errCheckName == nil {
		return c.Status(402).JSON(fiber.Map{
			"message": "category name already used.",
		})
	}

	if database.DB.Where("id = ?", id).Updates(&category).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Tidak dapat mengupdate data",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data berhasil diupdate",
		"data":    category,
	})
}

func CategoryDelete(c *fiber.Ctx) error {
	id := c.Params("id")

	var category entity.Category
	database.DB.First(&category, id)

	if category.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "not found",
		})
	}

	database.DB.Delete(&category)

	return c.JSON(fiber.Map{
		"message": "category deleted",
	})
}
