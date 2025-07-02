package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Dish struct {
	DishID              string `bson:"dish_id"`
	Name                string `bson:"name"`
	Quantity            int    `bson:"quantity"`
	SpecialInstructions string `bson:"special_instructions"`
	PreparationTime     string `bson:"preparation_time"`
}

type Order struct {
	ID                    primitive.ObjectID `bson:"_id,omitempty"`
	OrderID               string             `bson:"order_id"`
	TableID               string             `bson:"table_id"`
	CustomerID            string             `bson:"customer_id"`
	Status                string             `bson:"status"`
	OrderedAt             time.Time          `bson:"ordered_at"`
	TotalAmount           float64            `bson:"total_amount"`
	PaymentStatus         string             `bson:"payment_status"`
	PaymentMethod         string             `bson:"payment_method"`
	Dishes                []Dish             `bson:"dishes"`
	EstimatedDeliveryTime time.Time          `bson:"estimated_delivery_time"`
	DeliveredAt           time.Time          `bson:"delivered_at"`
	IsTakeaway            bool               `bson:"is_takeaway"`
	DiscountType          string             `bson:"discount_type"`
	OrderTaker            string             `bson:"order_taker"`
	Server                string             `bson:"server"`
}
