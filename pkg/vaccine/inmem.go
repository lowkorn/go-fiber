package vaccine

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/lowkorn/vaccine-reservation/util"
)

type InMemDriver struct {
	data     map[string]Vaccine
	uniquify map[string]struct{}
}

func (imd InMemDriver) Create(vc Vaccine) (Vaccine, error) {
	pID := strings.ToLower(vc.Name) + strings.ToLower(vc.Tel)
	uniqueKey := util.Hash(pID)
	if _, exist := imd.uniquify[uniqueKey]; exist {
		return Vaccine{}, errors.New("Can not insert duplications")
	}
	seed := time.Now().Unix()
	hashID := util.Hash(fmt.Sprint(seed))
	vc.ID = hashID
	imd.data[hashID] = vc
	imd.uniquify[uniqueKey] = struct{}{}
	return vc, nil
}

func (imd InMemDriver) GetAll() ([]Vaccine, error) {
	var vaccineReservation []Vaccine
	for _, vr := range imd.data {
		vaccineReservation = append(vaccineReservation, vr)
	}
	return vaccineReservation, nil
}

func (imd InMemDriver) GetByID(ID string) (Vaccine, error) {
	if vc, exist := imd.data[ID]; exist {
		return vc, nil
	}
	return Vaccine{}, nil
}

func (imd InMemDriver) Edit(ID string, vc Vaccine) (Vaccine, error) {
	if _, exist := imd.data[ID]; exist {
		vc.ID = ID
		imd.data[ID] = vc
		return vc, nil
	}
	return Vaccine{}, nil
}

func (imd InMemDriver) Delete(ID string) error {
	if _, exist := imd.data[ID]; exist {
		delete(imd.data, ID)
		return nil
	}
	return errors.New("Can not delete data")
}

func NewInmemInstance() InMemDriver {
	return InMemDriver{
		data:     map[string]Vaccine{},
		uniquify: map[string]struct{}{},
	}
}
