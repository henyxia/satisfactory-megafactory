package ui

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/henyxia/satisfactory-megafactory/internal/models"
)

func (u *UI) FactoryView(g *gin.Context) {
	factory := &models.Factory{}

	err := g.ShouldBindUri(factory)
	if err != nil {
		u.ErrNotFound(g)
		return
	}

	result := u.db.Preload("Floors").First(factory)

	if result.RowsAffected < 1 {
		u.ErrNotFound(g)
		return
	}

	var productionLineCount int64

	u.db.
		Joins("join floors on floors.factory_id = ?", factory.ID).
		Joins("join production_lines on production_lines.floor_id = floors.id").
		Count(&productionLineCount)

	g.HTML(http.StatusOK, "pages/factory-view.html", gin.H{
		"factory":             factory,
		"productionLineCount": productionLineCount,
	})
}
