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

type mgoOrdersRepository struct {
	baseRepo       *mongodb.BaseRepository
	collectionName string
}

func NewMgoOrdersRepository(client *mongo.Client) domain.OrdersRepository {
	collectionName := "orders"
	collection := client.Database(config.Cfgs.MgoDBName).Collection(collectionName)
	baseRepo := mongodb.NewBaseRepository(collection)
	m := &mgoOrdersRepository{
		baseRepo:       baseRepo,
		collectionName: collectionName,
	}
	return m
}

func (m *mgoOrdersRepository) Add(data *domain.OrderDTO) (*domain.OrderDTO, error) {
	insertResult, err := m.baseRepo.Add(data, 5*time.Second)
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	result := &domain.OrderDTO{
		Id:      insertResult.InsertedID.(primitive.ObjectID),
		Message: data.Message,
	}
	return result, err
}

func (m *mgoOrdersRepository) FindById(id string) (order *domain.OrderDTO, err error) {
	result, err := m.baseRepo.FindByID(id)
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	if err := result.Decode(&order); err != nil {
		glog.Error(err)
		return nil, err

	}
	return order, err
}

func (m *mgoOrdersRepository) UpdateById(id string, data *domain.OrderDTO) (*domain.OrderDTO, error) {
	update := bson.M{
		"$set": data,
	}
	result, err := m.baseRepo.UpdateByID(id, update, 5*time.Second)
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	var order *domain.OrderDTO
	if err := result.Decode(&order); err != nil {
		glog.Error(err)
		return nil, err

	}
	return order, err
}

func (m *mgoOrdersRepository) FindByIds(ids []string) ([]*domain.OrderDTO, error) {
	if len(ids) <= 0 {
		return []*domain.OrderDTO{}, nil
	}
	objIds, err := mongodb.ObjIds(ids)
	if err != nil {
		return nil, err
	}
	filter := bson.M{
		"_id": bson.M{
			"$in": objIds,
		},
	}
	return m.find(filter)
}

func (m *mgoOrdersRepository) find(filter bson.M, opts ...*options.FindOptions) ([]*domain.OrderDTO, error) {
	cur, _, err := m.baseRepo.Find(filter, 5*time.Second)
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	defer cur.Close(context.TODO())
	var results []*domain.OrderDTO
	for cur.Next(context.TODO()) {
		var result *domain.OrderDTO
		if err := cur.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return results, nil
}
