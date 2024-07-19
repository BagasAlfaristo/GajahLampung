package data

import (
	"gajahlampung/features/product"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName string
	Stock       uint
	Price       ProductPrice    `gorm:"foreignKey:ProductID"`
	Brand       ProductBrand    `gorm:"foreignKey:ProductID"`
	Category    ProductCategory `gorm:"foreignKey:ProductID"`
	Vehicle     VehicleDetail   `gorm:"foreignKey:ProductID"`
}

func GormToCore(gorm Product) product.Product {
	return product.Product{
		ID:          gorm.ID,
		ProductName: gorm.ProductName,
		Stock:       gorm.Stock,
		Price: product.ProductPrice{
			ID:           gorm.Price.ID,
			ProductID:    gorm.Price.ProductID,
			CostPrice:    gorm.Price.CostPrice,
			SellingPrice: gorm.Price.SellingPrice,
		},
		Brand: product.ProductBrand{
			ID:        gorm.Brand.ID,
			ProductID: gorm.Brand.ProductID,
			BrandName: gorm.Brand.BrandName,
		},
		Category: product.ProductCategory{
			ID:           gorm.Category.ID,
			ProductID:    gorm.Category.ProductID,
			CategoryName: gorm.Category.CategoryName,
		},
		Vehicle: product.VehicleDetail{
			ID:               gorm.Vehicle.ID,
			ProductID:        gorm.Vehicle.ProductID,
			VehicleBrandName: gorm.Vehicle.VehicleBrandName,
			VehicleName:      gorm.Vehicle.VehicleName,
			VehicleType:      gorm.Vehicle.VehicleType,
		},
	}
}
