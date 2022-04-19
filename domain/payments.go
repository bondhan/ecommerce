package domain

import "gorm.io/gorm"

type Cashiers struct {
	gorm.Model
	Name        string `gorm:"column:name"`
	Passcode    string `gorm:"column:passcode"`
	LoginStatus string `gorm:"column:login_status"`
}

func (m *Cashiers) TableName() string {
	return "cashiers"
}
