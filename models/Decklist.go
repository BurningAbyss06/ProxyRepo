package models

import "gorm.io/gorm"

type Decklist struct {
	gorm.Model

	Evento     string
	Top        int16
	Comentario string
	DeckId     uint
}
