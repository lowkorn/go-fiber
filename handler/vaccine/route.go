package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/lowkorn/vaccine-reservation/pkg/vaccine"
)

type VaccineReservationRoute struct {
	uc vaccine.IUsecase
}

func (r VaccineReservationRoute) GetReservation(c *fiber.Ctx) error {
	ID := c.Params("id")
	response, err := r.uc.GetVaccineReservationsByID(ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("fail to get reservation information")
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (r VaccineReservationRoute) CreateReservation(c *fiber.Ctx) error {
	var body vaccine.Vaccine
	err := c.BodyParser(&body)
	if err != nil {
		c.Status(fiber.ErrBadRequest.Code).SendString("invalid payload")
	}
	reponse, err := r.uc.MakeAVaccineReservations(body)
	if err != nil {
		c.Status(fiber.ErrBadRequest.Code).SendString("fail to create a reservations")
	}
	return c.Status(fiber.StatusOK).JSON(reponse)
}

func (r VaccineReservationRoute) UpdateReservation(c *fiber.Ctx) error {
	ID := c.Params("id")
	var body vaccine.Vaccine
	err := c.BodyParser(&body)
	if err != nil {
		c.Status(fiber.ErrBadRequest.Code).SendString("invalid payload")
	}
	response, err := r.uc.EditVaccineReservations(ID, body)
	if err != nil {
		c.Status(fiber.ErrBadRequest.Code).SendString("fail to create a reservations")
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (r VaccineReservationRoute) CancleReservation(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.Status(fiber.StatusOK).SendString(fmt.Sprintf("cancel reservations ID: %s", id))
}

func NewVaccineReservationRoute(vaccineUsecase vaccine.IUsecase) VaccineReservationRoute {
	return VaccineReservationRoute{}
}
