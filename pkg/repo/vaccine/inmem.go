package repo

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/lowkorn/vaccine-reservation/pkg/entity"
	"github.com/lowkorn/vaccine-reservation/pkg/util"
)

type VaccineInMem struct {
	data     map[string]entity.Vaccine
	uniquify map[string]struct{}
}

func (r VaccineInMem) Create(vc entity.Vaccine) (entity.Vaccine, error) {
	pID := strings.ToLower(vc.Name) + strings.ToLower(vc.Tel)
	uniqueKey := util.Hash(pID)
	if _, exist := r.uniquify[uniqueKey]; exist {
		return entity.Vaccine{}, errors.New("Can not insert duplications")
	}
	seed := time.Now().Unix()
	hashID := util.Hash(fmt.Sprint(seed))
	vc.ID = hashID
	r.data[hashID] = vc
	r.uniquify[uniqueKey] = struct{}{}
	return vc, nil
}

func (r VaccineInMem) GetAll() ([]entity.Vaccine, error) {
	var vaccineReservation []entity.Vaccine
	for _, vr := range r.data {
		vaccineReservation = append(vaccineReservation, vr)
	}
	return vaccineReservation, nil
}

func (r VaccineInMem) GetByID(ID string) (entity.Vaccine, error) {
	if vc, exist := r.data[ID]; exist {
		return vc, nil
	}
	return entity.Vaccine{}, nil
}

func (r VaccineInMem) Edit(ID string, vc entity.Vaccine) (entity.Vaccine, error) {
	if _, exist := r.data[ID]; exist {
		vc.ID = ID
		r.data[ID] = vc
		return vc, nil
	}
	return entity.Vaccine{}, nil
}

func (r VaccineInMem) Delete(ID string) error {
	if _, exist := r.data[ID]; exist {
		delete(r.data, ID)
		return nil
	}
	return errors.New("Can not delete data")
}

func NewVaccineInMem() VaccineInMem {
	return VaccineInMem{
		data:     map[string]entity.Vaccine{},
		uniquify: map[string]struct{}{},
	}
}
