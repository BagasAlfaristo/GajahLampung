package data

import "gorm.io/gorm"

type VehicleDetail struct {
	gorm.Model
	VehicleBrandName string
	VehicleName      string
	VehicleType      string
	ProductID        uint
}
