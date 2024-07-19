package handler

import (
	"gajahlampung/features/product"
)

type ProductRequest struct {
	ProductName      string  `json:"product_name" form:"product_name"`
	Stock            uint    `json:"stock" form:"stock"`
	CostPrice        float64 `json:"cost_price" form:"cost_price"`
	SellingPrice     float64 `json:"selling_price" form:"selling_price"`
	ProductID        uint    `json:"product_id" form:"product_id"`
	BrandName        string  `json:"brand_name" form:"brand_name"`
	CategoryName     string  `json:"category_name" form:"category_name"`
	VehicleBrandName string  `json:"vehicle_brand" form:"vehicle_brand"`
	VehicleName      string  `json:"vehicle_name" form:"vehicle_name"`
	VehicleType      string  `json:"vehicle_type" form:"vehicle_type"`
}

func RequestToCore(inputProduct ProductRequest) product.Product {
	return product.Product{
		ProductName: inputProduct.ProductName,
		Stock:       inputProduct.Stock,
		Price: product.ProductPrice{
			CostPrice:    inputProduct.CostPrice,
			SellingPrice: inputProduct.SellingPrice,
			ProductID:    inputProduct.ProductID,
		},
		Brand: product.ProductBrand{
			BrandName: inputProduct.BrandName,
			ProductID: inputProduct.ProductID,
		},
		Category: product.ProductCategory{
			CategoryName: inputProduct.CategoryName,
			ProductID:    inputProduct.ProductID,
		},
		Vehicle: product.VehicleDetail{
			VehicleBrandName: inputProduct.VehicleBrandName,
			VehicleName:      inputProduct.VehicleName,
			VehicleType:      inputProduct.VehicleType,
			ProductID:        inputProduct.ProductID,
		},
	}
}
