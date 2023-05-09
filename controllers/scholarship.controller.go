package controllers

import (
	"fmt"
	"go-miniproject-iqbalikhlasulamal/database"
	"go-miniproject-iqbalikhlasulamal/models/entity"
	"go-miniproject-iqbalikhlasulamal/models/request"
	"go-miniproject-iqbalikhlasulamal/models/response"
	"log"
	"sort"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func ScholarshipGetAll(c *fiber.Ctx) error {
	// get all scholarship data from database
	// name := c.Query("name")
	// categoryID := c.Query("category_id")
	var scholarships []entity.Scholarship

	result := database.DB.Debug().Find(
		&scholarships,
		// order by id desc
	).Order("id desc")

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "scholarship not found",
		})
	}

	// return response.ScholarshipResponse
	scholarshipResponse := []response.ScholarshipResponse{}

	for _, scholarship := range scholarships {
		scholarshipResponse = append(scholarshipResponse, response.ScholarshipResponse{
			ID:          scholarship.ID,
			Name:        scholarship.Name,
			Description: scholarship.Description,
			Deadline:    scholarship.Deadline,
			Link:        scholarship.Link,
			Thumbnail:   scholarship.Thumbnail,
			UserID:      scholarship.UserID,
			CategoryID:  scholarship.CategoryID,
			CreatedAt:   scholarship.CreatedAt,
		})
	}

	// sort by id desc
	sort.Slice(scholarshipResponse, func(i, j int) bool {
		return scholarshipResponse[i].ID < scholarshipResponse[j].ID
	})

	return c.JSON(scholarshipResponse)
}

func ScholarshipCreate(c *fiber.Ctx) error {
	scholarshipRequest := new(request.ScholarshipRequest)
	if err := c.BodyParser(scholarshipRequest); err != nil {
		return err
	}

	//	Validation Required Image
	var filenameString string

	filename := c.Locals("filename")
	log.Println("filename = ", filename)
	if filename == nil {
		return c.Status(422).JSON(fiber.Map{
			"message": "image cover is required.",
		})
	} else {
		filenameString = fmt.Sprintf("%v", filename)
	}

	var scholarship entity.Scholarship
	userInfo := c.Locals("userInfo").(jwt.MapClaims)
	userId := uint64(userInfo["id"].(float64))

	if err := c.BodyParser(&scholarship); err != nil {
		return err
	}

	newData := entity.Scholarship{
		UserID:      userId,
		CategoryID:  scholarship.CategoryID,
		Name:        scholarship.Name,
		Description: scholarship.Description,
		Thumbnail:   filenameString,
		Link:        scholarship.Link,
		Deadline:    scholarship.Deadline,
	}

	result := database.DB.Debug().Create(&newData)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to create scholarship",
			"error":   result.Error,
		})
	}

	scholarshipResponse := response.ScholarshipResponse{
		ID:          newData.ID,
		UserID:      newData.UserID,
		CategoryID:  newData.CategoryID,
		Name:        newData.Name,
		Description: newData.Description,
		Thumbnail:   newData.Thumbnail,
		Link:        newData.Link,
		Deadline:    newData.Deadline,
		CreatedAt:   newData.CreatedAt,
	}

	return c.JSON(fiber.Map{
		"message": "success create scholarship",
		"data":    scholarshipResponse,
	})

}

func ScholarshipUpdate(c *fiber.Ctx) error {
	scholarshipRequest := new(request.ScholarshipRequest)
	if err := c.BodyParser(scholarshipRequest); err != nil {
		return err
	}

	var scholarship entity.Scholarship
	userInfo := c.Locals("userInfo").(jwt.MapClaims)
	userId := uint64(userInfo["id"].(float64))

	if err := c.BodyParser(&scholarship); err != nil {
		return err
	}

	updatedData := entity.Scholarship{
		UserID:      userId,
		CategoryID:  scholarship.CategoryID,
		Name:        scholarship.Name,
		Description: scholarship.Description,
		Thumbnail:   scholarship.Thumbnail,
		Link:        scholarship.Link,
		Deadline:    scholarship.Deadline,
	}

	id := c.Params("id")

	var newData entity.Scholarship
	result := database.DB.Debug().Model(&newData).Where("id = ?", id).Updates(updatedData)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to update scholarship",
			"error":   result.Error,
		})
	}

	// return result after update
	result = database.DB.Debug().Where("id = ?", id).First(&newData)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to get scholarship",
			"error":   result.Error,
		})
	}

	scholarshipResponse := response.ScholarshipResponse{
		ID:          newData.ID,
		UserID:      newData.UserID,
		CategoryID:  newData.CategoryID,
		Name:        newData.Name,
		Description: newData.Description,
		Thumbnail:   newData.Thumbnail,
		Link:        newData.Link,
		Deadline:    newData.Deadline,
		CreatedAt:   newData.CreatedAt,
	}

	return c.JSON(fiber.Map{
		"message": "success update scholarship",
		"data":    scholarshipResponse,
	})
}

func ScholarshipDelete(c *fiber.Ctx) error {
	userInfo := c.Locals("userInfo").(jwt.MapClaims)
	userId := uint64(userInfo["id"].(float64))

	id := c.Params("id")

	scholarship := entity.Scholarship{}
	getScholarship := database.DB.Debug().Where("id = ?", id).First(&scholarship)
	if getScholarship.Error != nil {
		log.Println(getScholarship.Error)
	}

	if scholarship.UserID != userId {
		return c.Status(401).JSON(fiber.Map{
			"message": "failed",
			"error":   "unauthorized",
		})
	}

	result := database.DB.Debug().Delete(&scholarship)
	if result.Error != nil {
		log.Println(result.Error)
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func ScholarshipGetById(c *fiber.Ctx) error {
	scholarship := entity.Scholarship{}
	id := c.Params("id")

	result := database.DB.Debug().Where("id = ?", id).First(&scholarship)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "scholarship not found",
		})
	}

	scholarshipResponse := response.ScholarshipResponse{
		ID:          scholarship.ID,
		UserID:      scholarship.UserID,
		CategoryID:  scholarship.CategoryID,
		Name:        scholarship.Name,
		Description: scholarship.Description,
		Thumbnail:   scholarship.Thumbnail,
		Link:        scholarship.Link,
		Deadline:    scholarship.Deadline,
		CreatedAt:   scholarship.CreatedAt,
	}

	return c.JSON(scholarshipResponse)
}
