package dtos

import "time"

type Table struct {
	TableId      string       `json:"table_id"`
	NoOfSeats    int          `json:"no_of_seats"`
	ReservedBy   Reservation  `json:"reserved_by"`
	CurrentOrder OrderDetails `json:"current_order"`
}

type Reservation struct {
	Name         string    `json:"name"`
	Phone        string    `json:"phone"`
	ReservedFrom time.Time `json:"reserved_from"`
	ReservedTill time.Time `json:"reserved_till"`
}

type OrderDetails struct {
	OrderId    string `json:"order_id"`
	CustomerId string `json:"customer_id"`
}
