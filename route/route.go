package route

import (
	"go-miniproject-iqbalikhlasulamal/config"
	"go-miniproject-iqbalikhlasulamal/controllers"
	"go-miniproject-iqbalikhlasulamal/middleware"
	"go-miniproject-iqbalikhlasulamal/utils"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {

	route := r.Group("/api/v1")

	route.Static("/public", config.ProjectRootPath+"/public") // untuk alternative path jikalau path dibawah tidak terbaca
	route.Static("/public", "./public")                       // gunakan ini terlebih dahulu

	route.Post("/login", controllers.Login)

	route.Get("/user", middleware.Auth, controllers.UserGetAll)
	route.Get("/user/me", middleware.Auth, controllers.UserMe)
	route.Post("/user", controllers.UserCreate)
	route.Put("/user", middleware.Auth, controllers.UserUpdate)
	route.Put("/user/update-email", middleware.Auth, controllers.UserUpdateEmail)

	route.Get("/scholarship", controllers.ScholarshipGetAll)
	route.Post("/scholarship", middleware.Auth, utils.HandleSingleFile, controllers.ScholarshipCreate)
	route.Put("/scholarship/:id", middleware.Auth, controllers.ScholarshipUpdate)
	route.Delete("/scholarship/:id", middleware.Auth, controllers.ScholarshipDelete)
	route.Get("/scholarship/:id", controllers.ScholarshipGetById)

	route.Get("/category", controllers.CategoryGetAll)
	route.Get("/category/:id", controllers.CategoryGetByID)
	route.Post("/category", middleware.Auth, controllers.CategoryCreate)
	route.Put("/category/:id", middleware.Auth, controllers.CategoryUpdate)
	route.Delete("/category/:id", middleware.Auth, controllers.CategoryDelete)

	// r.Post("/book", utils.HandleSingleFile, controllers.BookCreate)

	// r.Post("/gallery", utils.HandleMultipleFile, controllers.PhotoCreate)
	// r.Delete("/gallery/:id", utils.HandleMultipleFile, controllers.PhotoDelete)
}
