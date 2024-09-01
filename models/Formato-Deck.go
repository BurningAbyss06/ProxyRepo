package models

import "gorm.io/gorm"

type F_D struct {
	gorm.Model

	IDFormato uint   `gorm:"not null;references:id"`
	IDDeck    uint   `gorm:"not null;references:id"`
	Region    string `gorm:"not null"`
}
