package controller

import (
	"github.com/anle/codebase/internal/model"
	"github.com/anle/codebase/internal/service"
	"github.com/anle/codebase/response"
	"github.com/gin-gonic/gin"
)

type CartController struct {
	cartService service.ICartService
}

func (cc *CartController) AddToCart(c *gin.Context) {
	var input model.AddToCartInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	result, _ := cc.cartService.AddToCart(c, input)

	response.HandleResult(c, result, nil)
}

func (cc *CartController) GetCart(c *gin.Context) {
	cart, result, _ := cc.cartService.GetCart(c)

	response.HandleResult(c, result, cart)
}

func (cc *CartController) DeleteCart(c *gin.Context) {
	var input model.DeleteCartInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	result, _ := cc.cartService.DeleteCartByID(c, input)

	response.HandleResult(c, result, nil)
}

func (cc *CartController) UpdateCart(c *gin.Context) {
	var input model.UpdateCartInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	result, _ := cc.cartService.UpdateCart(c, input)

	response.HandleResult(c, result, nil)
}

func NewCartController(cartService service.ICartService) *CartController {
	return &CartController{
		cartService: cartService,
	}
}
