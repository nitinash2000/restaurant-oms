package orders

import (
	"restaurant-oms/dtos"
	"restaurant-oms/models"
	"restaurant-oms/repository"
)

type OrderService interface {
	CreateOrder(req *dtos.Order) error
	UpdateOrder(id string, req *dtos.Order) error
	GetOrder(orderId string) (*dtos.Order, error)
	DeleteOrder(orderId string) error
}

type orderService struct {
	orderRepo repository.OrderRepo
}

func NewOrderService(orderRepo repository.OrderRepo) OrderService {
	return &orderService{
		orderRepo: orderRepo,
	}
}

func (o *orderService) CreateOrder(req *dtos.Order) error {
	orderModel := OrderDtosToModel(req)

	err := o.orderRepo.Create(orderModel)
	if err != nil {
		return err
	}

	return nil
}

func (o *orderService) UpdateOrder(id string, req *dtos.Order) error {
	orderModel := OrderDtosToModel(req)

	err := o.orderRepo.Update(id, orderModel)
	if err != nil {
		return err
	}

	return nil
}

func (o *orderService) GetOrder(orderId string) (*dtos.Order, error) {
	order, err := o.orderRepo.Get(orderId)
	if err != nil {
		return nil, err
	}

	result := OrderModelToDtos(order)

	return result, nil
}

func (o *orderService) DeleteOrder(orderId string) error {
	err := o.orderRepo.Delete(orderId)
	if err != nil {
		return err
	}

	return nil
}

func OrderModelToDtos(m *models.Order) *dtos.Order {
	var dishes []dtos.Dish
	for _, v := range m.Dishes {
		dishes = append(dishes, dtos.Dish{
			DishID:              v.DishID,
			Name:                v.Name,
			Quantity:            v.Quantity,
			SpecialInstructions: v.SpecialInstructions,
			PreparationTime:     v.PreparationTime,
		})
	}

	return &dtos.Order{
		OrderID:               m.OrderID,
		TableID:               m.TableID,
		CustomerID:            m.CustomerID,
		Status:                m.Status,
		OrderedAt:             m.OrderedAt,
		TotalAmount:           m.TotalAmount,
		PaymentStatus:         m.PaymentStatus,
		PaymentMethod:         m.PaymentMethod,
		Dishes:                dishes,
		EstimatedDeliveryTime: m.EstimatedDeliveryTime,
		DeliveredAt:           m.DeliveredAt,
		IsTakeaway:            m.IsTakeaway,
		DiscountType:          m.DiscountType,
		OrderTaker:            m.OrderTaker,
		Server:                m.Server,
	}
}

func OrderDtosToModel(m *dtos.Order) *models.Order {
	var dishModels []models.Dish
	for _, v := range m.Dishes {
		dishModels = append(dishModels, models.Dish{
			DishID:              v.DishID,
			Name:                v.Name,
			Quantity:            v.Quantity,
			SpecialInstructions: v.SpecialInstructions,
			PreparationTime:     v.PreparationTime,
		})
	}

	return &models.Order{
		OrderID:               m.OrderID,
		TableID:               m.TableID,
		CustomerID:            m.CustomerID,
		Status:                m.Status,
		OrderedAt:             m.OrderedAt,
		TotalAmount:           m.TotalAmount,
		PaymentStatus:         m.PaymentStatus,
		PaymentMethod:         m.PaymentMethod,
		Dishes:                dishModels,
		EstimatedDeliveryTime: m.EstimatedDeliveryTime,
		DeliveredAt:           m.DeliveredAt,
		IsTakeaway:            m.IsTakeaway,
		DiscountType:          m.DiscountType,
		OrderTaker:            m.OrderTaker,
		Server:                m.Server,
	}
}
