package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type OrderDTO struct {
	Id      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Message string             `json:"message,omitempty" bson:"message,omitempty"`
}

type ReqOrderData struct {
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type OrdersRepository interface {
	Add(data *OrderDTO) (*OrderDTO, error)
	FindById(id string) (*OrderDTO, error)
	FindByIds(ids []string) (orders []*OrderDTO, total int64, err error)
	UpdateById(id string, data *OrderDTO) (*OrderDTO, error)
}

type OrdersUseCase interface {
	UpdateById(id string, data *ReqOrderData) (*OrderDTO, *UCaseErr)
}
