package controllers

import "github.com/gofiber/fiber/v3"

type Controller interface {
	Handle(c fiber.Ctx) error
}