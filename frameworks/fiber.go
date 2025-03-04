package frameworks

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Fiber() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from Fiber")
	})

	fmt.Println("Fiber running on :3003")
	app.Listen(":3003")
}
