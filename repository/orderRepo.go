package repository

import (
	"context"
	"fmt"
	"restaurant-oms/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepo interface {
	Create(order *models.Order) error
	Update(orderId string, order *models.Order) error
	Get(orderId string) (*models.Order, error)
	Delete(orderId string) error
}

type orderRepo struct {
	collection *mongo.Collection
}

func NewOrderRepo(client *mongo.Client, database, collection string) OrderRepo {
	return &orderRepo{
		collection: client.Database(database).Collection(collection),
	}
}

func (r *orderRepo) Create(order *models.Order) error {
	_, err := r.collection.InsertOne(context.Background(), order)
	if err != nil {
		return fmt.Errorf("could not create order: %w", err)
	}
	return nil
}

func (r *orderRepo) Update(orderId string, order *models.Order) error {
	filter := bson.M{"order_id": orderId}
	update := bson.M{"$set": order}

	_, err := r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("could not update order: %w", err)
	}
	return nil
}

func (r *orderRepo) Get(orderId string) (*models.Order, error) {
	var order models.Order
	filter := bson.M{"order_id": orderId}
	err := r.collection.FindOne(context.Background(), filter).Decode(&order)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("order with id %s not found", orderId)
		}
		return nil, fmt.Errorf("could not retrieve order: %w", err)
	}
	return &order, nil
}

func (r *orderRepo) Delete(orderId string) error {
	filter := bson.M{"order_id": orderId}
	_, err := r.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("could not delete order: %w", err)
	}
	return nil
}
