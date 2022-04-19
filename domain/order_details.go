package domain

import "gorm.io/gorm"

type OrderDetails struct {
	gorm.Model
	Name             string `gorm:"column:name"`
	Price            int64  `gorm:"column:price"`
	Qty              int64  `gorm:"column:qty"`
	OrderID          int64  `gorm:"column:order_id"`
	ProductID        int64  `gorm:"column:product_id"`
	DiscountID       int64  `gorm:"column:discount_id"`
	TotalNormalPrice int64  `gorm:"column:total_normal_price"`
	TotalFinalPrice  int64  `gorm:"column:total_final_price"`
}

func (m *OrderDetails) TableName() string {
	return "order_details"
}
