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

func (*patientsUseCase) FindAll() ([]*domain.RespPatientData, *domain.UCaseErr) {
	return nil, nil
}

func (*patientsUseCase) FindById(id string) (*domain.RespPatientData, *domain.UCaseErr) {
	return nil, nil
}

func (*patientsUseCase) CreateOrderById(id string, data *domain.ReqOrderData) (*domain.OrderDTO, *domain.UCaseErr) {
	return nil, nil
}
