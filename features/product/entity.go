package product

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID          uint `gorm:"primaryKey"`
	ProductName string
	Stock       uint            `gorm:"column:stock"`
	Price       ProductPrice    `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Brand       ProductBrand    `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Category    ProductCategory `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Vehicle     VehicleDetail   `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type ProductPrice struct {
	gorm.Model
	ID           uint
	ProductID    uint
	CostPrice    float64
	SellingPrice float64
}

type ProductBrand struct {
	gorm.Model
	ID        uint
	ProductID uint
	BrandName string
}

type ProductCategory struct {
	gorm.Model
	ID           uint
	ProductID    uint
	CategoryName string
}

type VehicleDetail struct {
	gorm.Model
	ID               uint
	ProductID        uint
	VehicleBrandName string
	VehicleName      string
	VehicleType      string
}

type ProductModel interface {
	AddProduct(product Product) (Product, error)
	GetProduct() ([]Product, error)
	GetProductByName(productName string) ([]Product, error)
}

type ProductService interface {
	AddProduct(Product Product) (Product, error)
	GetProduct() ([]Product, error)
	GetProductByName(productName string) ([]Product, error)
}
