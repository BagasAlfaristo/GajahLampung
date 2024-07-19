package data

import "gorm.io/gorm"

type ProductPrice struct {
	gorm.Model
	CostPrice    float64
	SellingPrice float64
	ProductID    uint
}
