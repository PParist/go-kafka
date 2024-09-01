package entities

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	AccountUUID   string  `gorm:"type:char(36);uniqueIndex;not null"`
	AccountHolder string  `gorm:"type:varchar(100);not null"`
	AccountType   string  `gorm:"type:varchar(50);not null"`
	Balance       float64 `gorm:"type:decimal(18,2);default:0.0;not null"`
}
