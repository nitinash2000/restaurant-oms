package dtos

import "time"

type Dish struct {
	DishID              string `json:"dish_id"`
	Name                string `json:"name"`
	Quantity            int    `json:"quantity"`
	SpecialInstructions string `json:"special_instructions"`
	PreparationTime     string `json:"preparation_time"`
}

type Order struct {
	OrderID               string    `json:"order_id"`
	TableID               string    `json:"table_id"`
	CustomerID            string    `json:"customer_id"`
	Status                string    `json:"status"`
	OrderedAt             time.Time `json:"ordered_at"`
	TotalAmount           float64   `json:"total_amount"`
	PaymentStatus         string    `json:"payment_status"`
	PaymentMethod         string    `json:"payment_method"`
	Dishes                []Dish    `json:"dishes"`
	EstimatedDeliveryTime time.Time `json:"estimated_delivery_time"`
	DeliveredAt           time.Time `json:"delivered_at"`
	IsTakeaway            bool      `json:"is_takeaway"`
	DiscountType          string    `json:"discount_type"`
	OrderTaker            string    `json:"order_taker"`
	Server                string    `json:"server"`
}
