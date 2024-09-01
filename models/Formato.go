package models

import (
	"time"

	"gorm.io/gorm"
)

type Formato struct {
	gorm.Model

	Nombre        string    `gorm:"not null; unique_index"`
	InicioFormato time.Time `gorm:"not null"`
	FinFormato    time.Time `gorm:"not null"`
	SetsFormato   string    `gorm:"not null; unique_index"`
}
