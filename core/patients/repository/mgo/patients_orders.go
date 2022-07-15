package mgo

import (
	"jubo-go-api/config"
	"jubo-go-api/domain"
	"jubo-go-api/infrastructures/mongodb"

	"go.mongodb.org/mongo-driver/mongo"
)

type mgoPatientsRepository struct {
	baseRepo       *mongodb.BaseRepository
	collectionName string
}

func NewMgoPatientsRepository(
	client *mongo.Client,
) domain.PatientsRepository {
	collectionName := "patients"
	collection := client.Database(config.Cfgs.MgoDBName).Collection(collectionName)
	baseRepo := mongodb.NewBaseRepository(collection)
	m := &mgoPatientsRepository{
		baseRepo:       baseRepo,
		collectionName: collectionName,
	}
	return m
}

func (m *mgoPatientsRepository) Add(data *domain.PatientDTO) (*domain.PatientDTO, error) {
	return nil, nil
}

func (m *mgoPatientsRepository) FindAll() ([]*domain.PatientDTO, error) {
	return nil, nil
}

func (m *mgoPatientsRepository) FindById(id string) (*domain.PatientDTO, error) {
	return nil, nil
}

func (m *mgoPatientsRepository) UpdateById(id string, data *domain.PatientDTO) (*domain.PatientDTO, error) {
	return nil, nil
}
