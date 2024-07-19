package migrations

import (
	_dataProduct "gajahlampung/features/product/data"

	"gorm.io/gorm"
)

func InitMigrations(db *gorm.DB) {
	db.AutoMigrate(&_dataProduct.ProductPrice{})
	db.AutoMigrate(&_dataProduct.Product{})
	db.AutoMigrate(&_dataProduct.ProductBrand{})
	db.AutoMigrate(&_dataProduct.VehicleDetail{})
	db.AutoMigrate(&_dataProduct.ProductCategory{})
}
