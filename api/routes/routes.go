package routes

import "github.com/gofiber/fiber/v2"

var App *fiber.App

func InicializeRoutes() {
	InicializeRoutesIVS()
}
