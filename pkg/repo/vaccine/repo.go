package repo

import "github.com/lowkorn/vaccine-reservation/pkg/entity"

type IVaccine interface {
	GetAll() ([]entity.Vaccine, error)
	GetByID(ID string) (entity.Vaccine, error)
	Create(vc entity.Vaccine) (entity.Vaccine, error)
	Edit(ID string, vc entity.Vaccine) (entity.Vaccine, error)
	Delete(ID string) error
}
