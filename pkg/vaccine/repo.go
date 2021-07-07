package vaccine

type IRepo interface {
	GetAll() ([]Vaccine, error)
	GetByID(ID string) (Vaccine, error)
	Create(vc Vaccine) (Vaccine, error)
	Edit(ID string, vc Vaccine) (Vaccine, error)
	Delete(ID string) error
}
