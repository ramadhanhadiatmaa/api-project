package middlewares

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Auth(ctx *fiber.Ctx) error{
 if err := godotenv.Load(); err != nil {
	log.Fatalf("Error loading .env file %v", err)
 }

 secret_key := os.Getenv("SECRET_KEY")

 header := ctx.Get("apikey")

 if header == "" || header != secret_key {
	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "Unauthorized Access Request",
	})
 }
 return ctx.Next()
}