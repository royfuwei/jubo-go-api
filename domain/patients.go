package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type PatientDTO struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	OrderIds []string           `json:"orderIds" bson:"orderIds,omitempty"`
}

type RespPatientData struct {
	*PatientDTO `json:",inline"`
	Orders      []*OrderDTO `json:"orders,omitempty"`
}

type PatientsRepository interface {
	Add(data *PatientDTO) (*PatientDTO, error)
	FindAll() ([]*PatientDTO, error)
	FindById(id string) (*PatientDTO, error)
	UpdateById(id string, data *PatientDTO) (*PatientDTO, error)
}

type PatientUseCase interface {
	FindAll() ([]*RespPatientData, *UCaseErr)
	FindById(id string) (*RespPatientData, *UCaseErr)
	CreateOrderById(id string, data *ReqOrderData) (*OrderDTO, *UCaseErr)
}
