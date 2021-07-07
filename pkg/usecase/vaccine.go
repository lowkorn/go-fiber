package usecase

import (
	"github.com/lowkorn/vaccine-reservation/pkg/entity"
	repo "github.com/lowkorn/vaccine-reservation/pkg/repo/vaccine"
)

type IUsecase interface {
	MakeAReservation(vc entity.Vaccine) (entity.Vaccine, error)
	GetReservationByID(ID string) (entity.Vaccine, error)
	GetAllReservation() ([]entity.Vaccine, error)
	EditReservation(ID string, vc entity.Vaccine) (entity.Vaccine, error)
	CancelReservation(ID string) error
}

type Vaccine struct {
	Repo repo.IVaccine
}

func (uc Vaccine) MakeAReservation(vc entity.Vaccine) (entity.Vaccine, error) {
	response, err := uc.Repo.Create(vc)
	return response, err
}

func (uc Vaccine) GetAllReservation() ([]entity.Vaccine, error) {
	return uc.Repo.GetAll()
}

func (uc Vaccine) GetReservationByID(ID string) (entity.Vaccine, error) {
	response, err := uc.Repo.GetByID(ID)
	return response, err
}

func (uc Vaccine) EditReservation(ID string, vc entity.Vaccine) (entity.Vaccine, error) {
	response, err := uc.Repo.Edit(ID, vc)
	return response, err
}

func (uc Vaccine) CancelReservation(ID string) error {
	return uc.Repo.Delete(ID)
}

func NewVaccine(vaccineRepo repo.IVaccine) Vaccine {
	return Vaccine{
		Repo: vaccineRepo,
	}
}
