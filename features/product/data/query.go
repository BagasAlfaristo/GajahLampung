package data

import (
	"gajahlampung/features/product"

	"gorm.io/gorm"
)

type productModel struct {
	db *gorm.DB
}

func New(db *gorm.DB) product.ProductModel {
	return &productModel{
		db: db,
	}
}

func (pm *productModel) AddProduct(product product.Product) (product.Product, error) {
	tx := pm.db.Create(&product)
	if tx.Error != nil {
		return product, tx.Error
	}
	return product, nil
}

func (pm *productModel) GetProduct() ([]product.Product, error) {
	var productGorm []Product
	if err := pm.db.Preload("Price").Preload("Brand").Preload("Category").Preload("Vehicle").Find(&productGorm).Error; err != nil {
		return nil, err
	}
	var productCore []product.Product
	for _, v := range productGorm {
		productCore = append(productCore, GormToCore(v))
	}
	return productCore, nil
}

func (pm *productModel) GetProductByName(productName string) ([]product.Product, error) {
	var productGorm []Product
	tx := pm.db.Preload("Price").Preload("Brand").Preload("Category").Preload("Vehicle").Where("vehicle_name LIKE ?", "%"+productName+"%").Find(&productGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var productCore []product.Product
	for _, v := range productGorm {
		productCore = append(productCore, GormToCore(v))
	}
	return productCore, nil
}
