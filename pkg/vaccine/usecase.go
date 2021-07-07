package vaccine

type IUsecase interface {
	MakeAReservation(vc Vaccine) (Vaccine, error)
	GetReservationByID(ID string) (Vaccine, error)
	GetAllReservation() ([]Vaccine, error)
	EditReservation(ID string, vc Vaccine) (Vaccine, error)
	CancelReservation(ID string) error
}

type Usecase struct {
	Repo IRepo
}

func (uc Usecase) MakeAReservation(vc Vaccine) (Vaccine, error) {
	response, err := uc.Repo.Create(vc)
	return response, err
}

func (uc Usecase) GetAllReservation() ([]Vaccine, error) {
	return uc.Repo.GetAll()
}

func (uc Usecase) GetReservationByID(ID string) (Vaccine, error) {
	response, err := uc.Repo.GetByID(ID)
	return response, err
}

func (uc Usecase) EditReservation(ID string, vc Vaccine) (Vaccine, error) {
	response, err := uc.Repo.Edit(ID, vc)
	return response, err
}

func (uc Usecase) CancelReservation(ID string) error {
	return uc.Repo.Delete(ID)
}

func NewUsecase(vaccineRepo IRepo) Usecase {
	return Usecase{
		Repo: vaccineRepo,
	}
}
