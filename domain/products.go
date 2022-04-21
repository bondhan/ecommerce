package domain

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	Name       string `gorm:"column:name"`
	Image      string `gorm:"column:image"`
	Stock      int64  `gorm:"column:stock"`
	Price      int64  `gorm:"column:price"`
	CategoryID uint   `gorm:"column:category_id"`
	DiscountID *uint  `gorm:"column:discount_id"`
}

func (m *Products) TableName() string {
	return "products"
}
