package vaccine

type IRepo interface {
	GetByID(ID string) (Vaccine, error)
	Create(vc Vaccine) (Vaccine, error)
	Edit(ID string, vc Vaccine) (Vaccine, error)
}
