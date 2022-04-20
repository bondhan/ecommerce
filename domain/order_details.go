package domain

import "gorm.io/gorm"

type OrderDetails struct {
	gorm.Model
	Name             string  `gorm:"column:name"`
	Price            int64   `gorm:"column:price"`
	Qty              int64   `gorm:"column:qty"`
	OrderID          uint    `gorm:"column:order_id"`
	ProductID        uint    `gorm:"column:product_id"`
	DiscountID       *uint   `gorm:"column:discount_id"`
	TotalNormalPrice float64 `gorm:"column:total_normal_price"`
	TotalFinalPrice  float64 `gorm:"column:total_final_price"`
}

func (m *OrderDetails) TableName() string {
	return "order_details"
}
