package vaccine

type IUsecase interface {
	MakeAVaccineReservations(vc Vaccine) (Vaccine, error)
	GetVaccineReservationsByID(ID string) (Vaccine, error)
	EditVaccineReservations(ID string, vc Vaccine) (Vaccine, error)
}

type Usecase struct {
	Repo IRepo
}

func (uc Usecase) MakeAVaccineReservations(vc Vaccine) (Vaccine, error) {
	response, err := uc.Repo.Create(vc)
	return response, err
}

func (uc Usecase) GetVaccineReservationsByID(ID string) (Vaccine, error) {
	response, err := uc.Repo.GetByID(ID)
	return response, err
}

func (uc Usecase) EditVaccineReservations(ID string, vc Vaccine) (Vaccine, error) {
	response, err := uc.Repo.Edit(ID, vc)
	return response, err
}

func NewUsecase(vaccineRepo IRepo) Usecase {
	return Usecase{
		Repo: vaccineRepo,
	}
}
