package domain

import (
	"gorm.io/gorm"
	"time"
)

type Discount struct {
	gorm.Model
	Qty       int64     `gorm:"column:qty"`
	Type      string    `gorm:"column:type"`
	Result    int64     `gorm:"column:result"`
	ExpiredAt time.Time `gorm:"column:expired_at"`
}

func (m *Discount) TableName() string {
	return "discounts"
}
