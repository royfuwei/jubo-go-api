package usecases

import "jubo-go-api/domain"

type patientsUseCase struct {
	patientsRepo domain.PatientsRepository
	ordersRepo   domain.OrdersRepository
}

func NewPatientsUseCase(
	patientsRepo domain.PatientsRepository,
	ordersRepo domain.OrdersRepository,
) domain.PatientUseCase {
	return &patientsUseCase{
		patientsRepo: patientsRepo,
		ordersRepo:   ordersRepo,
	}
}

func (ucase *patientsUseCase) AddOne(data *domain.ReqAddOne) (*domain.RespPatientData, *domain.UCaseErr) {
	return nil, nil
}

func (ucase *patientsUseCase) FindAll() ([]*domain.RespPatientData, *domain.UCaseErr) {
	return nil, nil
}

func (ucase *patientsUseCase) FindById(id string) (*domain.RespPatientData, *domain.UCaseErr) {
	return nil, nil
}

func (ucase *patientsUseCase) AddOrderById(id string, data *domain.ReqOrderData) (*domain.RespPatientData, *domain.UCaseErr) {
	return nil, nil
}
