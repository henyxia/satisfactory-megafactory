package ui

import (
	"gorm.io/gorm"
)

type UI struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UI {
	return &UI{
		db: db,
	}
}
