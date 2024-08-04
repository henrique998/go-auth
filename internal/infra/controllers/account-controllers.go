package controllers

import "github.com/gofiber/fiber/v3"

type CreateAccountController interface {
	Handle(c fiber.Ctx) error
}
