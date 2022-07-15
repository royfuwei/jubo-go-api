package usecases

import "jubo-go-api/domain"

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
	return result, nil
}
