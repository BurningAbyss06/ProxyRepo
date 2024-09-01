package models

import "gorm.io/gorm"

type Deck struct {
	gorm.Model

	Nombre   string     `gorm:"not null; unique_index"`
	Decklist []Decklist `gorm:"references:id"`
}
