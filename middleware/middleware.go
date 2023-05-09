package middleware

import (
	"go-miniproject-iqbalikhlasulamal/utils"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Auth(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims, err := utils.DecodeToken(token)
	if err != nil {
		log.Println("err :: ", err)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	// log.Println("claims :: ", claims)

	ctx.Locals("userInfo", claims)

	return ctx.Next()
}
