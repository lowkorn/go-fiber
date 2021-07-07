package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/lowkorn/vaccine-reservation/pkg/vaccine"
)

type VaccineReservationRoute struct {
	uc vaccine.IUsecase
}

func (r VaccineReservationRoute) GetReservation(c *fiber.Ctx) error {
	ID := c.Params("id")
	response, err := r.uc.GetReservationByID(ID)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("fail to get reservation information")
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (r VaccineReservationRoute) GetAllVaccineReservation(c *fiber.Ctx) error {
	reponse, err := r.uc.GetAllReservation()
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("fail to get reservation informations")
	}
	return c.Status(fiber.StatusOK).JSON(reponse)
}

func (r VaccineReservationRoute) CreateReservation(c *fiber.Ctx) error {
	var body vaccine.Vaccine
	err := c.BodyParser(&body)
	if err != nil {
		log.Println(err, body)
		return c.Status(fiber.ErrBadRequest.Code).SendString("invalid payload")
	}
	reponse, err := r.uc.MakeAReservation(body)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.ErrBadRequest.Code).SendString("fail to create a reservations")
	}
	return c.Status(fiber.StatusOK).JSON(reponse)
}

func (r VaccineReservationRoute) UpdateReservation(c *fiber.Ctx) error {
	ID := c.Params("id")
	var body vaccine.Vaccine
	err := c.BodyParser(&body)
	if err != nil {
		log.Println(err, body)
		return c.Status(fiber.ErrBadRequest.Code).SendString("invalid payload")
	}
	response, err := r.uc.EditReservation(ID, body)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.ErrBadRequest.Code).SendString("fail to create a reservations")
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (r VaccineReservationRoute) CancleReservation(c *fiber.Ctx) error {
	ID := c.Params("id")
	err := r.uc.CancelReservation(ID)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.ErrBadRequest.Code).SendString("fail to cancel a reservations")
	}
	return c.Status(fiber.StatusOK).SendString("success")
}

func NewVaccineReservationRoute(vaccineUsecase vaccine.IUsecase) VaccineReservationRoute {
	return VaccineReservationRoute{
		uc: vaccineUsecase,
	}
}
