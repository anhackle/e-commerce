package controller

import (
	"github.com/anle/codebase/internal/model"
	"github.com/anle/codebase/internal/service"
	"github.com/anle/codebase/response"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService service.IProductService
}

func (pc *ProductController) DeleteProduct(c *gin.Context) {
	var input model.DeleteProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	result, _ := pc.productService.DeleteProduct(c, input)
	response.HandleResult(c, result, nil)
}

func (pc *ProductController) UpdateProduct(c *gin.Context) {
	var input model.UpdateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	result, _ := pc.productService.UpdateProduct(c, input)
	response.HandleResult(c, result, nil)
}

func (pc *ProductController) GetProducts(c *gin.Context) {
	var input model.GetProductsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	products, result, _ := pc.productService.GetProducts(c, input)
	response.HandleResult(c, result, products)
}
func (pc *ProductController) CreateProduct(c *gin.Context) {
	var input model.CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	result, _ := pc.productService.CreateProduct(c, input)
	response.HandleResult(c, result, nil)
}

func (pc *ProductController) GetProductsForAdmin(c *gin.Context) {
	var input model.GetProductsForAdminInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	products, result, _ := pc.productService.GetProductsForAdmin(c, input)

	response.HandleResult(c, result, products)
}

func NewProductController(productService service.IProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}
