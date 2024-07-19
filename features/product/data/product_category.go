package data

import "gorm.io/gorm"

type ProductCategory struct {
	gorm.Model
	CategoryName string
	ProductID    uint
}
