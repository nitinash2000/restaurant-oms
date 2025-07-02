package tables

import (
	"restaurant-oms/dtos"
	"restaurant-oms/models"
	"restaurant-oms/repository"
)

type TableService interface {
	CreateTable(req *dtos.Table) error
	UpdateTable(id string, req *dtos.Table) error
	GetTable(tableId string) (*dtos.Table, error)
	DeleteTable(tableId string) error
}

type tableService struct {
	tableRepo repository.TableRepo
}

func NewTableService(tableRepo repository.TableRepo) TableService {
	return &tableService{
		tableRepo: tableRepo,
	}
}

func (o *tableService) CreateTable(req *dtos.Table) error {
	tableModel := TableDtosToModel(req)

	err := o.tableRepo.Create(tableModel)
	if err != nil {
		return err
	}

	return nil
}

func (o *tableService) UpdateTable(id string, req *dtos.Table) error {
	tableModel := TableDtosToModel(req)

	err := o.tableRepo.Update(id, tableModel)
	if err != nil {
		return err
	}

	return nil
}

func (o *tableService) GetTable(tableId string) (*dtos.Table, error) {
	table, err := o.tableRepo.Get(tableId)
	if err != nil {
		return nil, err
	}

	result := TableModelToDtos(table)

	return result, nil
}

func (o *tableService) DeleteTable(tableId string) error {
	err := o.tableRepo.Delete(tableId)
	if err != nil {
		return err
	}

	return nil
}

func TableModelToDtos(m *models.Table) *dtos.Table {
	table := &dtos.Table{
		TableId:   m.TableId,
		NoOfSeats: m.NoOfSeats,
		ReservedBy: dtos.Reservation{
			Name:         m.ReservedBy.Name,
			Phone:        m.ReservedBy.Phone,
			ReservedFrom: m.ReservedBy.ReservedFrom,
			ReservedTill: m.ReservedBy.ReservedTill,
		},
		CurrentOrder: dtos.OrderDetails{
			OrderId:    m.CurrentOrder.OrderId,
			CustomerId: m.CurrentOrder.CustomerId,
		},
	}

	return table
}

func TableDtosToModel(m *dtos.Table) *models.Table {
	return &models.Table{
		TableId:   m.TableId,
		NoOfSeats: m.NoOfSeats,
		ReservedBy: models.Reservation{
			Name:         m.ReservedBy.Name,
			Phone:        m.ReservedBy.Phone,
			ReservedFrom: m.ReservedBy.ReservedFrom,
			ReservedTill: m.ReservedBy.ReservedTill,
		},
		CurrentOrder: models.OrderDetails{
			OrderId:    m.CurrentOrder.OrderId,
			CustomerId: m.CurrentOrder.CustomerId,
		},
	}
}
