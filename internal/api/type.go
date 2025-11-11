package api

import (
	"gorm.io/gorm"
)

type API struct {
	db *gorm.DB
}

func New(db *gorm.DB) *API {
	return &API{
		db: db,
	}
}
