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

type RespFindAll struct {
	Total int64              `json:"total"`
	Items []*RespPatientData `json:"items"`
}

type PatientsRepository interface {
	Add(data *PatientDTO) (*PatientDTO, error)
	FindAll() (patients []*PatientDTO, total int64, err error)
	FindById(id string) (*PatientDTO, error)
	UpdateById(id string, data *PatientDTO) (*PatientDTO, error)
}

type PatientUseCase interface {
	AddOne(data *PatientDTO) (*RespPatientData, *UCaseErr)
	FindAll() ([]*RespPatientData, *UCaseErr)
	FindById(id string) (*RespPatientData, *UCaseErr)
	AddOrderById(id string, data *ReqOrderData) (*RespPatientData, *UCaseErr)
}
