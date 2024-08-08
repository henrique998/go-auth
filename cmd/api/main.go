package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"github.com/henrique998/go-auth/internal/configs/logger"
	"github.com/henrique998/go-auth/internal/infra/database"
	"github.com/henrique998/go-auth/internal/infra/endpoints"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	db := database.ConnectToDb()
	defer db.Close()

	app := fiber.New()

	app.Use(limiter.New(limiter.Config{
		Expiration: 30 * time.Second,
		Max:        5,
		LimitReached: func(c fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too many requests, please try again later.",
			})
		},
	}))

	endpoints.SetupAccountsEndpoints(app, db)
	logger.Error("Project startup", app.Listen(":3333"))
}
