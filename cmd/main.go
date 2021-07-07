package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	handler "github.com/lowkorn/vaccine-reservation/pkg/handler/vaccine"
	repo "github.com/lowkorn/vaccine-reservation/pkg/repo/vaccine"
	"github.com/lowkorn/vaccine-reservation/pkg/service"
	"github.com/lowkorn/vaccine-reservation/pkg/usecase"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	mongoClient, err := service.NewMongoConnection("mongodb://root:root@localhost:27017")
	if err != nil {
		log.Panicln(err)
	}

	// vacReservRepo := repo.NewVaccineInMem()
	vacReservRepo := repo.NewVaccineMongo(&mongoClient)
	vacReservUC := usecase.NewVaccine(vacReservRepo)
	vacReserveRoute := handler.NewVaccineReservationRoute(vacReservUC)

	app.Get("/health", func(fiberC *fiber.Ctx) error {
		message := map[string]string{
			"status": "ok",
		}
		return fiberC.Status(fiber.StatusOK).JSON(message)
	})

	app.Get("vaccine-reservations/:id", vacReserveRoute.GetReservation)
	app.Get("vaccine-reservations", vacReserveRoute.GetAllVaccineReservation)
	app.Post("vaccine-reservations", vacReserveRoute.CreateReservation)
	app.Put("vaccine-reservations/:id", vacReserveRoute.UpdateReservation)
	app.Delete("vaccine-reservations/:id", vacReserveRoute.CancleReservation)

	app.Listen(":5000")
}
