package models

import (
	"gorm.io/gorm"
)

type Weather struct {
	gorm.Model
	LocationName string
	Temp         float32
}
