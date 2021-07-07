package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	handler "github.com/lowkorn/vaccine-reservation/handler/vaccine"
	"github.com/lowkorn/vaccine-reservation/pkg/vaccine"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	vacReservRepo := vaccine.NewInmemInstance()
	vacReservUC := vaccine.NewUsecase(vacReservRepo)
	vacReserveRoute := handler.NewVaccineReservationRoute(vacReservUC)

	app.Get("/health", func(fiberC *fiber.Ctx) error {
		message := map[string]string{
			"status": "ok",
		}
		return fiberC.Status(fiber.StatusOK).JSON(message)
	})

	app.Get("vaccine-reservations/:id", vacReserveRoute.GetReservation)
	app.Post("vaccine-reservations", vacReserveRoute.CreateReservation)
	app.Put("vaccine-reservations/:id", vacReserveRoute.UpdateReservation)
	app.Delete("vaccine-reservations/:id", vacReserveRoute.GetReservation)

	app.Listen(":5000")
}
