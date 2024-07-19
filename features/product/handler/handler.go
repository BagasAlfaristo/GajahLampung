package handler

import (
	"fmt"
	"gajahlampung/features/product"
	"gajahlampung/utils/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	ProductService product.ProductService
}

func New(ps product.ProductService) *ProductHandler {
	return &ProductHandler{
		ProductService: ps,
	}
}

func (ph *ProductHandler) AddProduct(c echo.Context) error {
	var newProductRequest ProductRequest
	errBind := c.Bind(&newProductRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("Error bind data"+errBind.Error(), nil))
	}

	inputCore := RequestToCore(newProductRequest)

	fmt.Println()
	_, errInsert := ph.ProductService.AddProduct(inputCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(errInsert.Error(), nil))
	}
	return c.JSON(http.StatusCreated, responses.JSONWebResponse("Add product successful!", nil))
}

func (ph *ProductHandler) GetProduct(c echo.Context) error {
	products, err := ph.ProductService.GetProduct()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Error retrieving products: "+err.Error(), nil))
	}

	var response []ProductResponse
	for _, v := range products {
		response = append(response, AllGormToCore(v))
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("Product retrive successfully", response))
}

func (ph *ProductHandler) GetProductByName(c echo.Context) error {
	name := c.QueryParam("name")
	products, err := ph.ProductService.GetProductByName(name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("Error retrieving products: "+err.Error(), nil))
	}

	var response []ProductResponse
	for _, v := range products {
		response = append(response, AllGormToCore(v))
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("Product retrive successfully", response))
}
