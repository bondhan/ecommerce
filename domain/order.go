package domain

import "gorm.io/gorm"

type Orders struct {
	gorm.Model
	TotalPrice    int64  `gorm:"column:total_price"`
	TotalPaid     int64  `gorm:"column:total_paid"`
	TotalReturn   int64  `gorm:"column:total_return"`
	CashierID     uint   `gorm:"column:cashier_id"`
	PaymentTypeID uint   `gorm:"column:payment_type_id"`
	InvoicePDF    string `gorm:"column:invoice_pdf"`
	Downloaded    int    `gorm:"column:downloaded"`
}

func (m *Orders) TableName() string {
	return "orders"
}
