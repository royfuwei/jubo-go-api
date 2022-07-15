package mgo

import (
	"context"
	"jubo-go-api/config"
	"jubo-go-api/domain"
	"jubo-go-api/infrastructures/mongodb"
	"time"

	"github.com/golang/glog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	result, err := m.baseRepo.Add(data, 5*time.Second)
	if err != nil {
		return nil, err
	}
	patient := &domain.PatientDTO{
		Id:       result.InsertedID.(primitive.ObjectID),
		Name:     data.Name,
		OrderIds: data.OrderIds,
	}
	return patient, nil
}

func (m *mgoPatientsRepository) FindAll() ([]*domain.PatientDTO, int64, error) {
	filter := bson.M{}
	return m.find(filter)
}

func (m *mgoPatientsRepository) FindById(id string) (*domain.PatientDTO, error) {
	result, err := m.baseRepo.FindByID(id)
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	var patient *domain.PatientDTO
	if err := result.Decode(&patient); err != nil {
		glog.Error(err)
		return nil, err
	}
	return patient, nil
}

func (m *mgoPatientsRepository) UpdateById(id string, data *domain.PatientDTO) (*domain.PatientDTO, error) {
	update := bson.M{
		"$set": data,
	}
	result, err := m.baseRepo.UpdateByID(id, update, 5*time.Second)
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	var patient *domain.PatientDTO
	if err := result.Decode(&patient); err != nil {
		glog.Error(err)
		return nil, err
	}
	return patient, nil
}

func (m *mgoPatientsRepository) find(filter bson.M, opts ...*options.FindOptions) (patients []*domain.PatientDTO, total int64, err error) {
	cur, total, err := m.baseRepo.Find(filter, 5*time.Second)
	if err != nil {
		return nil, 0, err
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var result *domain.PatientDTO
		if err = cur.Decode(&result); err != nil {
			return nil, total, err
		}
		patients = append(patients, result)
	}
	if err = cur.Err(); err != nil {
		return nil, total, err
	}
	return patients, total, nil
}
