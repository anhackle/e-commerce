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

func (oc *OrderController) GetOrder(c *gin.Context) {
	panic("unimplemented")
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

func (oc *OrderController) DeleteOrder(c *gin.Context) {
	panic("unimplemented")
}

func NewOrderController(orderService service.IOrderService) *OrderController {
	return &OrderController{
		orderService: orderService,
	}
}
