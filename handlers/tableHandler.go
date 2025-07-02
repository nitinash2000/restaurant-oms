package handlers

import (
	"net/http"
	"restaurant-oms/dtos"
	"restaurant-oms/services/tables"

	"github.com/gin-gonic/gin"
)

type tableHandler struct {
	tableService tables.TableService
}

func NewTableHandler(tableService tables.TableService) *tableHandler {
	return &tableHandler{
		tableService: tableService,
	}
}

func (o *tableHandler) GetTable(ctx *gin.Context) {
	id := ctx.Param("id")

	table, err := o.tableService.GetTable(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, table)
}

func (o *tableHandler) CreateTable(ctx *gin.Context) {
	var req *dtos.Table

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = o.tableService.CreateTable(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Table created successfully"})
}

func (o *tableHandler) DeleteTable(ctx *gin.Context) {
	id := ctx.Param("id")

	err := o.tableService.DeleteTable(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Table deleted successfully"})
}

func (o *tableHandler) UpdateTable(ctx *gin.Context) {
	id := ctx.Param("id")

	var req dtos.Table
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = o.tableService.UpdateTable(id, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Updated table successfully"})
}
