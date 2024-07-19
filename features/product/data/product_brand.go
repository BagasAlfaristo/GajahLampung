package data

import "gorm.io/gorm"

type ProductBrand struct {
	gorm.Model
	BrandName string
	ProductID uint
}
