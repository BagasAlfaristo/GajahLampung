package handler

import "gajahlampung/features/product"

type ProductResponse struct {
	ID          uint            `json:"id"`
	ProductName string          `json:"product_name"`
	Stock       uint            `json:"stock"`
	Price       ProductPrice    `json:"price"`
	Brand       ProductBrand    `json:"brand"`
	Category    ProductCategory `json:"type"`
	Vehicle     VehicleDetail   `json:"vehicle_details"`
}

type ProductPrice struct {
	ProductID    uint    `json:"product_id"`
	CostPrice    float64 `json:"cost_price"`
	SellingPrice float64 `json:"selling_price"`
}

type ProductBrand struct {
	ProductID uint   `json:"product_id"`
	BrandName string `json:"brand_name"`
}

type ProductCategory struct {
	ProductID    uint   `json:"product_id"`
	CategoryName string `json:"category_name"`
}

type VehicleDetail struct {
	ProductID        uint   `json:"product_id"`
	VehicleBrandName string `json:"vehicle_brand"`
	VehicleName      string `json:"vehicle_name"`
	VehicleType      string `json:"vehicle_type"`
}

func AllGormToCore(inputProduct product.Product) ProductResponse {
	return ProductResponse{
		ID:          inputProduct.ID,
		ProductName: inputProduct.ProductName,
		Stock:       inputProduct.Stock,
		Price: ProductPrice{
			ProductID:    inputProduct.Price.ProductID,
			CostPrice:    inputProduct.Price.CostPrice,
			SellingPrice: inputProduct.Price.SellingPrice,
		},
		Brand: ProductBrand{
			ProductID: inputProduct.Brand.ProductID,
			BrandName: inputProduct.Brand.BrandName,
		},
		Category: ProductCategory{
			ProductID:    inputProduct.Category.ProductID,
			CategoryName: inputProduct.Category.CategoryName,
		},
		Vehicle: VehicleDetail{
			ProductID:        inputProduct.Vehicle.ProductID,
			VehicleBrandName: inputProduct.Vehicle.VehicleBrandName,
			VehicleName:      inputProduct.Vehicle.VehicleName,
			VehicleType:      inputProduct.Vehicle.VehicleType,
		},
	}
}
