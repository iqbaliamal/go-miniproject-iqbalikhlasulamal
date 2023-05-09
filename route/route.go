package route

import (
	"go-miniproject-iqbalikhlasulamal/config"
	"go-miniproject-iqbalikhlasulamal/controllers"
	"go-miniproject-iqbalikhlasulamal/middleware"
	"go-miniproject-iqbalikhlasulamal/utils"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {

	r.Group("/api/v1")

	r.Static("/public", config.ProjectRootPath+"/public") // untuk alternative path jikalau path dibawah tidak terbaca
	r.Static("/public", "./public")                       // gunakan ini terlebih dahulu

	r.Post("/login", controllers.Login)

	r.Get("/user", middleware.Auth, controllers.UserGetAll)
	r.Get("/user/me", middleware.Auth, controllers.UserMe)
	r.Post("/user", controllers.UserCreate)
	r.Put("/user", middleware.Auth, controllers.UserUpdate)
	r.Put("/user/update-email", middleware.Auth, controllers.UserUpdateEmail)

	r.Get("/scholarship", controllers.ScholarshipGetAll)
	r.Post("/scholarship", middleware.Auth, utils.HandleSingleFile, controllers.ScholarshipCreate)
	r.Put("/scholarship/:id", middleware.Auth, controllers.ScholarshipUpdate)
	r.Delete("/scholarship/:id", middleware.Auth, controllers.ScholarshipDelete)
	r.Get("/scholarship/:id", controllers.ScholarshipGetById)

	r.Get("/category", controllers.CategoryGetAll)
	r.Get("/category/:id", controllers.CategoryGetByID)
	r.Post("/category", middleware.Auth, controllers.CategoryCreate)
	r.Put("/category/:id", middleware.Auth, controllers.CategoryUpdate)
	r.Delete("/category/:id", middleware.Auth, controllers.CategoryDelete)

	// r.Post("/book", utils.HandleSingleFile, controllers.BookCreate)

	// r.Post("/gallery", utils.HandleMultipleFile, controllers.PhotoCreate)
	// r.Delete("/gallery/:id", utils.HandleMultipleFile, controllers.PhotoDelete)
}
