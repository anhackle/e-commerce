package controller

import (
	"github.com/anle/codebase/internal/model"
	"github.com/anle/codebase/internal/service"
	"github.com/anle/codebase/response"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService service.IOrderService
}

func (oc *OrderController) GetOrders(c *gin.Context) {
	var input model.GetOrdersInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	orders, result, _ := oc.orderService.GetOrders(c, input)

	response.HandleResult(c, result, orders)
}

func (oc *OrderController) CreateOrder(c *gin.Context) {
	var input model.CreateOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	result, _ := oc.orderService.CreateOrder(c, input)

	response.HandleResult(c, result, nil)
}

func (oc *OrderController) GetOrder(c *gin.Context) {
	var input model.GetOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	orderDetail, result, _ := oc.orderService.GetOrder(c, input)

	response.HandleResult(c, result, orderDetail)
}

func (oc *OrderController) UpdateStatus(c *gin.Context) {
	var input model.UpdateStatusInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	result, _ := oc.orderService.UpdateStatus(c, input)

	response.HandleResult(c, result, nil)
}

func (oc *OrderController) GetOrdersForAdmin(c *gin.Context) {
	var input model.GetOrdersForAdminInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	orders, result, _ := oc.orderService.GetOrdersForAdmin(c, input)

	response.HandleResult(c, result, orders)
}

func NewOrderController(orderService service.IOrderService) *OrderController {
	return &OrderController{
		orderService: orderService,
	}
}
