package api

import (
	"fmt"
	"golang-ivs/api/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

var port string

func Run() {
	port = os.Getenv("api_port")

	if len(port) < 1 {
		port = "5000"
	}

	fmt.Printf(" Aplicação iniciada (em: http://localhost:%s)\n", port)
	log.Fatal(routes.App.Listen(fmt.Sprintf(":%v", port)))
}

func init() {

	routes.App = fiber.New()

	routes.App.Use(func(c *fiber.Ctx) error {
		c.Type("json", "utf-8")
		return c.Next()
	})

	routes.App.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(
			struct {
				Response string `json:"mensagem"`
			}{
				Response: "api started ;)",
			})
	})

	routes.InicializeRoutes()

}
