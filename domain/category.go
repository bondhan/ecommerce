package domain

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `gorm:"column:name"`
}

func (m *Category) TableName() string {
	return "categories"
}
