package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name         string        `gorm:"type:varchar(50);not null" json:"name"`
	Description  string        `gorm:"type:varchar(200)" json:"description"`
	UserID       uint          `gorm:"not null" json:"userId"`
	User         User          `gorm:"foreignKey:UserID" json:"-"`
	Transactions []Transaction `gorm:"foreignKey:CategoryID" json:"transactions,omitempty"`
}

func (Category) TableName() string {
	return "categories"
}
