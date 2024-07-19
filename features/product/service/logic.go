package service

import (
	"fmt"
	"gajahlampung/features/product"
)

type ProductService struct {
	ProductModel product.ProductModel
}

func New(pm product.ProductModel) product.ProductService {
	return &ProductService{
		ProductModel: pm,
	}
}

func (pm *ProductService) AddProduct(productCore product.Product) (product.Product, error) {
	fmt.Println("service", productCore)
	return pm.ProductModel.AddProduct(productCore)
}

func (pm *ProductService) GetProduct() ([]product.Product, error) {
	return pm.ProductModel.GetProduct()
}

func (pm *ProductService) GetProductByName(productName string) ([]product.Product, error) {
	return pm.ProductModel.GetProductByName(productName)
}
