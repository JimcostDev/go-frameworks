package frameworks

import (
	"log"

	"github.com/JimcostDev/go-frameworks/routes"
	"github.com/gofiber/fiber/v2"
)

func Fiber() {
	app := fiber.New()

	// Configura las rutas
	routes.SetupRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello world"})
	})

	port := ":3003"
	log.Printf("Servidor corriendo en el puerto: %s", port)
	log.Fatal(app.Listen(port))
}
