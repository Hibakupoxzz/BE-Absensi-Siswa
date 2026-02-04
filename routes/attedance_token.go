package routes

import (
	"github.com/KicauOrgspark/BE-Absensi-Siswa/handlers"
	"github.com/KicauOrgspark/BE-Absensi-Siswa/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRouteAttedanceToken(api fiber.Router) {
	token := api.Group("/token")

	token.Get("/create", middleware.AdminRoute, handlers.Create)
}