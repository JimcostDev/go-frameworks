package routes

import (
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes registra las rutas principales de la aplicación.
func SetupRoutes(app *fiber.App) {
	// Agrupamos las rutas de la API bajo un prefijo (opcional)
	api := app.Group("/api")
	SetupProductRoutes(api)
}
