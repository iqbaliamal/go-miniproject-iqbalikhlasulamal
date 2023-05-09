package test

import (
	"bytes"
	"encoding/json"
	"go-miniproject-iqbalikhlasulamal/controllers"
	"go-miniproject-iqbalikhlasulamal/database"
	"go-miniproject-iqbalikhlasulamal/middleware"
	"go-miniproject-iqbalikhlasulamal/models/entity"
	"go-miniproject-iqbalikhlasulamal/models/request"
	"go-miniproject-iqbalikhlasulamal/utils"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("Error loading .env file")
	}
	database.DatabaseInit()

	app := fiber.New()

	// Add Route for testing
	app.Post("/api/v1/user", controllers.UserCreate)

	// Test register success
	user := entity.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password123",
	}

	payload, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodPost, "/api/v1/user", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, resp.StatusCode)

}

func TestLogin(t *testing.T) {
	// Setup
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("Error loading .env file")
	}
	database.DatabaseInit()
	app := fiber.New()
	// Routes
	app.Post("/api/v1/login", controllers.Login)
	// Test
	loginRequest := request.LoginRequest{
		Email:    "test@mail.com",
		Password: "password",
	}
	jsonLogin, _ := json.Marshal(loginRequest)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewBuffer(jsonLogin))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		log.Fatal(err)
	}

	// set timeout for response
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)
}

func TestCategoryGetAll(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("Error loading .env file")
	}
	database.DatabaseInit()

	app := fiber.New()
	// Routes
	app.Get("/api/v1/category", controllers.CategoryGetAll)
	// Test
	req := httptest.NewRequest(http.MethodGet, "/api/v1/category", nil)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, 200, resp.StatusCode)
}

func TestCategoryGetByID(t *testing.T) {
	// Setup
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("Error loading .env file")
	}
	database.DatabaseInit()
	app := fiber.New()
	// Routes
	app.Get("/api/v1/category/:id", controllers.CategoryGetByID)
	// Test
	req := httptest.NewRequest(http.MethodGet, "/api/v1/category/3", nil)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, 200, resp.StatusCode)
}

func TestCategoryCreate(t *testing.T) {
	// Setup
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("Error loading .env file")
	}
	database.DatabaseInit()
	app := fiber.New()
	// Routes
	app.Post("/api/v1/category", controllers.CategoryCreate)
	// Test
	categoryRequest := request.CategoryRequest{
		Name: "Category 1122",
	}
	jsonCategory, _ := json.Marshal(categoryRequest)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/category", bytes.NewBuffer(jsonCategory))
	req.Header.Set("Content-Type", "application/json")
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAbWFpbC5jb20iLCJleHAiOjE2ODM2MjMwMTEsImlkIjoxLCJuYW1lIjoidGVzdCIsInJvbGUiOiJ1c2VyIn0.R1cYy5W6qWO0J-Jf3IZ2VdeiBgCYoN5mEItcPSawqFM"
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := app.Test(req)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, 200, resp.StatusCode)
}

func TestCategoryUpdate(t *testing.T) {
	// Setup
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("Error loading .env file")
	}
	database.DatabaseInit()
	app := fiber.New()
	// Routes
	app.Put("/api/v1/category/:id", controllers.CategoryUpdate)
	// Test
	categoryRequest := request.CategoryRequest{
		Name: "Category 12344",
	}
	jsonCategory, _ := json.Marshal(categoryRequest)
	req := httptest.NewRequest(http.MethodPut, "/api/v1/category/2", bytes.NewBuffer(jsonCategory))
	req.Header.Set("Content-Type", "application/json")
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAbWFpbC5jb20iLCJleHAiOjE2ODM2MjMwMTEsImlkIjoxLCJuYW1lIjoidGVzdCIsInJvbGUiOiJ1c2VyIn0.R1cYy5W6qWO0J-Jf3IZ2VdeiBgCYoN5mEItcPSawqFM"
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := app.Test(req)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, 200, resp.StatusCode)
}

func TestCategoryDelete(t *testing.T) {
	// Setup
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("Error loading .env file")
	}
	database.DatabaseInit()
	app := fiber.New()
	// Routes
	app.Delete("/api/v1/category/:id", controllers.CategoryDelete)
	// Test
	req := httptest.NewRequest(http.MethodDelete, "/api/v1/category/2", nil)
	req.Header.Set("Content-Type", "application/json")
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAbWFpbC5jb20iLCJleHAiOjE2ODM2MjMwMTEsImlkIjoxLCJuYW1lIjoidGVzdCIsInJvbGUiOiJ1c2VyIn0.R1cYy5W6qWO0J-Jf3IZ2VdeiBgCYoN5mEItcPSawqFM"
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := app.Test(req)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUserMe(t *testing.T) {
	// Setup
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("Error loading .env file")
	}
	database.DatabaseInit()

	app := fiber.New()
	// Routes
	app.Get("/api/v1/user/me", middleware.Auth, controllers.UserMe)
	// Test
	req := httptest.NewRequest(http.MethodGet, "/api/v1/user/me", nil)
	req.Header.Set("Content-Type", "application/json")
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAbWFpbC5jb20iLCJleHAiOjE2ODM2MjMwMTEsImlkIjoxLCJuYW1lIjoidGVzdCIsInJvbGUiOiJ1c2VyIn0.R1cYy5W6qWO0J-Jf3IZ2VdeiBgCYoN5mEItcPSawqFM"
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := app.Test(req)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUserUpdate(t *testing.T) {
	// Setup
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("Error loading .env file")
	}
	database.DatabaseInit()
	app := fiber.New()
	// Routes
	app.Put("/api/v1/user", middleware.Auth, controllers.UserUpdate)
	// Test
	userRequest := request.UserUpdateRequest{
		Name: "test update",
	}
	jsonUser, _ := json.Marshal(userRequest)
	req := httptest.NewRequest(http.MethodPut, "/api/v1/user", bytes.NewBuffer(jsonUser))
	req.Header.Set("Content-Type", "application/json")
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAbWFpbC5jb20iLCJleHAiOjE2ODM2MjMwMTEsImlkIjoxLCJuYW1lIjoidGVzdCIsInJvbGUiOiJ1c2VyIn0.R1cYy5W6qWO0J-Jf3IZ2VdeiBgCYoN5mEItcPSawqFM"
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := app.Test(req)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, 200, resp.StatusCode)
}

func TestScholarshipGetAll(t *testing.T) {
	// Setup
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("Error loading .env file")
	}
	database.DatabaseInit()
	app := fiber.New()
	// Routes
	app.Get("/api/v1/scholarship", controllers.ScholarshipGetAll)
	// Test
	req := httptest.NewRequest(http.MethodGet, "/api/v1/scholarship", nil)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, 200, resp.StatusCode)
}

func TestScholarshipGetByID(t *testing.T) {
	// Setup
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("Error loading .env file")
	}
	database.DatabaseInit()
	app := fiber.New()
	// Routes
	app.Get("/api/v1/scholarship/:id", controllers.ScholarshipGetById)
	// Test
	req := httptest.NewRequest(http.MethodGet, "/api/v1/scholarship/12", nil)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, 200, resp.StatusCode)
}

func TestScholarshipCreate(t *testing.T) {
	// Setup
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("Error loading .env file")
	}
	database.DatabaseInit()
	app := fiber.New()
	// Routes
	app.Post("/api/v1/scholarship", middleware.Auth, utils.HandleSingleFile, controllers.ScholarshipCreate)
	// Test
	scholarshipRequest := request.ScholarshipRequest{
		Name:        "Scholarship 1 123",
		Description: "Scholarship 1",
		Link:        "https://www.google.com",
		Deadline:    "2021-01-01",
		CategoryID:  1,
	}
	jsonScholarship, _ := json.Marshal(scholarshipRequest)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/scholarship", bytes.NewBuffer(jsonScholarship))
	req.Header.Set("Content-Type", "application/json")
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAbWFpbC5jb20iLCJleHAiOjE2ODM2MjMwMTEsImlkIjoxLCJuYW1lIjoidGVzdCIsInJvbGUiOiJ1c2VyIn0.R1cYy5W6qWO0J-Jf3IZ2VdeiBgCYoN5mEItcPSawqFM"
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := app.Test(req)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 200, resp.StatusCode)
}

func TestScholarshipUpdate(t *testing.T) {
	// Setup
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("Error loading .env file")
	}
	database.DatabaseInit()
	app := fiber.New()
	// Routes
	app.Put("/api/v1/scholarship/:id", middleware.Auth, utils.HandleSingleFile, controllers.ScholarshipUpdate)
	// Test
	scholarshipRequest := request.ScholarshipRequest{
		Name:        "Scholarship 1",
		Description: "Scholarship 1",
		CategoryID:  1,
	}
	jsonScholarship, _ := json.Marshal(scholarshipRequest)
	req := httptest.NewRequest(http.MethodPut, "/api/v1/scholarship/12", bytes.NewBuffer(jsonScholarship))
	req.Header.Set("Content-Type", "application/json")
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAbWFpbC5jb20iLCJleHAiOjE2ODM2MjMwMTEsImlkIjoxLCJuYW1lIjoidGVzdCIsInJvbGUiOiJ1c2VyIn0.R1cYy5W6qWO0J-Jf3IZ2VdeiBgCYoN5mEItcPSawqFM"
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := app.Test(req)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 200, resp.StatusCode)
}

func TestScholarshipDelete(t *testing.T) {
	// Setup
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("Error loading .env file")
	}
	database.DatabaseInit()
	app := fiber.New()
	// Routes
	app.Delete("/api/v1/scholarship/:id", middleware.Auth, controllers.ScholarshipDelete)
	// Test
	req := httptest.NewRequest(http.MethodDelete, "/api/v1/scholarship/13", nil)
	req.Header.Set("Content-Type", "application/json")
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAbWFpbC5jb20iLCJleHAiOjE2ODM2MjMwMTEsImlkIjoxLCJuYW1lIjoidGVzdCIsInJvbGUiOiJ1c2VyIn0.R1cYy5W6qWO0J-Jf3IZ2VdeiBgCYoN5mEItcPSawqFM"
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := app.Test(req)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 200, resp.StatusCode)
}
