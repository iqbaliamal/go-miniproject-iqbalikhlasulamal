package main

import (
	"go-miniproject-iqbalikhlasulamal/database"
	"go-miniproject-iqbalikhlasulamal/database/migration"
	"go-miniproject-iqbalikhlasulamal/route"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}
	//	INITIAL DATABASE
	database.DatabaseInit()
	//	RUN MIGRATION
	migration.RunMigration()

	app := fiber.New()

	// INITIAL ROUTE
	route.RouteInit(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	errListen := app.Listen(":" + port)
	if errListen != nil {
		log.Println("Fail to listen go fiber server")
		os.Exit(1)
	}
}
