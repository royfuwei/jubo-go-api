package usecases

import (
	"jubo-go-api/domain"
	"jubo-go-api/domain/category"
	"jubo-go-api/domain/errcode"
	"jubo-go-api/infrastructures/tools"
)

type ordersUseCase struct {
	ordersRepo   domain.OrdersRepository
	patientsRepo domain.PatientsRepository
}

func NewOrdersUseCase(
	ordersRepo domain.OrdersRepository,
	patientsRepo domain.PatientsRepository,
) domain.OrdersUseCase {
	return &ordersUseCase{
		ordersRepo:   ordersRepo,
		patientsRepo: patientsRepo,
	}
}

func (ucase *ordersUseCase) UpdateById(id string, data *domain.ReqOrderData) (*domain.OrderDTO, *domain.UCaseErr) {
	var result *domain.OrderDTO
	update := &domain.OrderDTO{
		Message: data.Message,
	}
	result, err := ucase.ordersRepo.UpdateById(id, update)
	if err != nil {
		return nil, tools.NewUCaseErr(category.Orders, errcode.Default, err, nil)
	}
	return result, nil
}

func (ucase *ordersUseCase) FindManyByPatientId(patientId string) (*domain.RespOrders, *domain.UCaseErr) {
	patient, err := ucase.patientsRepo.FindById(patientId)
	if err != nil {
		return nil, tools.NewUCaseErr(category.Orders, errcode.Default, err, nil)
	}
	results, total, err := ucase.ordersRepo.FindByIds(patient.OrderIds)
	if err != nil {
		return nil, tools.NewUCaseErr(category.Orders, errcode.Default, err, nil)
	}
	return &domain.RespOrders{
		Total: total,
		Items: results,
	}, nil
}
