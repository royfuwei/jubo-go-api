package usecases

import (
	"fmt"
	"jubo-go-api/domain"
	"jubo-go-api/domain/category"
	"jubo-go-api/domain/errcode"
	"jubo-go-api/infrastructures/tools"
)

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

func (ucase *patientsUseCase) getRespPatientData(data *domain.PatientDTO) (*domain.RespPatientData, *domain.UCaseErr) {
	orderIds := data.OrderIds
	orders, _, err := ucase.ordersRepo.FindByIds(orderIds)
	if err != nil {
		return nil, tools.NewUCaseErr(category.Patients, errcode.Default, err, nil)
	}
	return &domain.RespPatientData{
		PatientDTO: data,
		Orders:     orders,
	}, nil
}

func (ucase *patientsUseCase) getRespPatientDataList(data []*domain.PatientDTO) ([]*domain.RespPatientData, *domain.UCaseErr) {
	results := []*domain.RespPatientData{}
	mixOrderIds := []string{}
	for _, patient := range data {
		orderIds := patient.OrderIds
		mixOrderIds = append(mixOrderIds, orderIds...)
	}
	orders, _, err := ucase.ordersRepo.FindByIds(mixOrderIds)
	if err != nil {
		return results, tools.NewUCaseErr(category.Patients, errcode.Default, err, nil)
	}
	ordersMap := make(map[string]*domain.OrderDTO)
	for _, order := range orders {
		ordersMap[order.Id.Hex()] = order
	}
	for _, patient := range data {
		orderIds := patient.OrderIds
		orders := []*domain.OrderDTO{}
		for _, orderId := range orderIds {
			orders = append(orders, ordersMap[orderId])
		}

		result := &domain.RespPatientData{
			PatientDTO: patient,
			Orders:     orders,
		}

		results = append(results, result)
	}
	return results, nil
}

func (ucase *patientsUseCase) AddOne(data *domain.ReqAddOne) (*domain.RespPatientData, *domain.UCaseErr) {
	create := &domain.PatientDTO{
		Name:     data.Name,
		OrderIds: []string{},
	}
	patient, err := ucase.patientsRepo.Add(create)
	if err != nil {
		return nil, tools.NewUCaseErr(category.Patients, errcode.Default, err, nil)
	}
	return ucase.getRespPatientData(patient)
}

func (ucase *patientsUseCase) FindAll() (*domain.RespFindAll, *domain.UCaseErr) {
	patients, total, err := ucase.patientsRepo.FindAll()
	if err != nil {
		return nil, tools.NewUCaseErr(category.Patients, errcode.Default, err, nil)
	}
	respPatientDataList, uCaseErr := ucase.getRespPatientDataList(patients)
	if uCaseErr != nil {
		return nil, uCaseErr
	}
	return &domain.RespFindAll{
		Items: respPatientDataList,
		Total: total,
	}, nil
}

func (ucase *patientsUseCase) FindById(id string) (*domain.RespPatientData, *domain.UCaseErr) {
	patient, err := ucase.patientsRepo.FindById(id)
	if err != nil {
		return nil, tools.NewUCaseErr(category.Patients, errcode.Default, err, nil)
	}
	return ucase.getRespPatientData(patient)
}

func (ucase *patientsUseCase) AddOrderById(id string, data *domain.ReqOrderData) (*domain.RespPatientData, *domain.UCaseErr) {
	patient, err := ucase.patientsRepo.FindById(id)
	if err != nil {
		return nil, tools.NewUCaseErr(category.Patients, errcode.Default, err, nil)
	}
	addOrder := &domain.OrderDTO{
		Message: data.Message,
	}
	order, err := ucase.ordersRepo.Add(addOrder)
	if err != nil {
		return nil, tools.NewUCaseErr(category.Patients, errcode.Default, err, nil)
	}
	orderIds := patient.OrderIds
	fmt.Printf("orderIds: %v \n", orderIds)
	orderIds = append(orderIds, order.Id.Hex())
	update := &domain.PatientDTO{
		OrderIds: orderIds,
	}
	ucase.patientsRepo.UpdateById(id, update)

	return ucase.FindById(id)
}
