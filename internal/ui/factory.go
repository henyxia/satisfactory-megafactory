package ui

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/henyxia/satisfactory-megafactory/internal/models"
)

func (u *UI) FactoryList(g *gin.Context) {
	var factories []models.Factory

	u.db.Find(&factories)

	g.HTML(http.StatusOK, "pages/factory-list.html", gin.H{
		"factories": factories,
	})
}

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

	var buildings []models.Building
	u.db.Order("name").Find(&buildings)

	response := gin.H{
		"factory":             factory,
		"productionLineCount": productionLineCount,
		"buildings":           buildings,
	}

	hpFloor := &models.Floor{}
	result = u.db.Where("factory_id = ? and type = 1", factory.ID).Order("level asc").First(hpFloor)
	if result.RowsAffected > 0 {
		response["homepageFloor"] = hpFloor.ID
	}

	g.HTML(http.StatusOK, "pages/factory-view.html", response)
}
