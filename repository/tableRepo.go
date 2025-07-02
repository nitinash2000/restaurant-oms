package repository

import (
	"context"
	"fmt"
	"restaurant-oms/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TableRepo interface {
	Create(table *models.Table) error
	Update(tableId string, table *models.Table) error
	Get(tableId string) (*models.Table, error)
	Delete(tableId string) error
}

type tableRepo struct {
	collection *mongo.Collection
}

func NewTableRepo(client *mongo.Client, database, collection string) TableRepo {
	return &tableRepo{
		collection: client.Database(database).Collection(collection),
	}
}

func (r *tableRepo) Create(table *models.Table) error {
	_, err := r.collection.InsertOne(context.Background(), table)
	if err != nil {
		return fmt.Errorf("could not create table: %w", err)
	}
	return nil
}

func (r *tableRepo) Update(tableId string, table *models.Table) error {
	filter := bson.M{"table_id": tableId}
	update := bson.M{"$set": table}

	_, err := r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("could not update table: %w", err)
	}
	return nil
}

func (r *tableRepo) Get(tableId string) (*models.Table, error) {
	var table models.Table
	filter := bson.M{"table_id": tableId}
	err := r.collection.FindOne(context.Background(), filter).Decode(&table)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("table with id %s not found", tableId)
		}
		return nil, fmt.Errorf("could not retrieve table: %w", err)
	}
	return &table, nil
}

func (r *tableRepo) Delete(tableId string) error {
	filter := bson.M{"table_id": tableId}
	_, err := r.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("could not delete table: %w", err)
	}
	return nil
}
