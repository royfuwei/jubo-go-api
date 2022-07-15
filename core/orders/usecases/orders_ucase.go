package usecases

import (
	"jubo-go-api/domain"
	"jubo-go-api/domain/category"
	"jubo-go-api/domain/errcode"
	"jubo-go-api/infrastructures/tools"
)

type ordersUseCase struct {
	ordersRepo domain.OrdersRepository
}

func NewOrdersUseCase(
	ordersRepo domain.OrdersRepository,
) domain.OrdersUseCase {
	return &ordersUseCase{
		ordersRepo: ordersRepo,
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
