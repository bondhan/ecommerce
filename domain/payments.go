package domain

import "gorm.io/gorm"

type Payments struct {
	gorm.Model
	Name string  `gorm:"column:name"`
	Type string  `gorm:"column:type"`
	Logo *string `gorm:"column:logo"`
}

func (m *Payments) TableName() string {
	return "payment_types"
}
