package models

import "time"

type Table struct {
	TableId      string       `bson:"table_id"`
	NoOfSeats    int          `bson:"no_of_seats"`
	ReservedBy   Reservation  `bson:"reserved_by"`
	CurrentOrder OrderDetails `bson:"current_order"`
}

type Reservation struct {
	Name         string    `bson:"name"`
	Phone        string    `bson:"phone"`
	ReservedFrom time.Time `bson:"reserved_from"`
	ReservedTill time.Time `bson:"reserved_till"`
}

type OrderDetails struct {
	OrderId    string `bson:"order_id"`
	CustomerId string `bson:"customer_id"`
}
