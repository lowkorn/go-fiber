package vaccine

import (
	"errors"
	"fmt"
	"strings"

	"github.com/lowkorn/vaccine-reservation/util"
)

type InMemDriver struct {
	Data map[string]Vaccine
}

func (imd InMemDriver) Create(vc Vaccine) (Vaccine, error) {
	ID := strings.ToLower(vc.Name) + strings.ToLower(vc.Tel)
	hashID := util.Hash(ID)
	if _, exist := imd.Data[hashID]; exist {
		return Vaccine{}, errors.New("Can not insert duplications")
	}
	vc.ID = hashID
	imd.Data[hashID] = vc
	return vc, nil
}

func (imd InMemDriver) GetByID(ID string) (Vaccine, error) {
	fmt.Println("ID", ID)
	fmt.Println("ID", ID)
	if vc, exist := imd.Data[ID]; exist {
		return vc, nil
	}
	return Vaccine{}, errors.New("Not exist")
}

func (imd InMemDriver) Edit(ID string, vc Vaccine) (Vaccine, error) {
	if _, exist := imd.Data[ID]; exist {
		imd.Data[ID] = vc
		return vc, nil
	}
	return Vaccine{}, errors.New("Con not edit in-exist data")
}

func NewInmemInstance() InMemDriver {
	return InMemDriver{
		Data: map[string]Vaccine{},
	}
}
