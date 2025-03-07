package routes

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// SetupProductRoutes configura las rutas de productos
func SetupProductRoutes(router fiber.Router) {
	productRouter := router.Group("/products")
	productRouter.Get("/", getProducts)
	productRouter.Get("/:id", getProductByID)
}

// getProducts maneja la solicitud de lista de productos
func getProducts(c *fiber.Ctx) error {
	size, err := strconv.Atoi(c.Query("size", "10"))
	if err != nil {
		size = 10
	}
	products := make([]map[string]interface{}, size)
	for i := 1; i <= size; i++ {
		products[i-1] = map[string]interface{}{
			"id":   i,
			"name": "Product " + strconv.Itoa(i),
		}
	}
	return c.JSON(products)
}

// getProductByID obtiene un producto por su ID
func getProductByID(c *fiber.Ctx) error {
	id := c.Params("id")

	return c.JSON(fiber.Map{
		"id":   id,
		"name": "Product " + id,
	})
}
