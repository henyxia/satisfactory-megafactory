package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/henyxia/satisfactory-megafactory/internal/models"
)

func (a *API) BuildingRecipe(g *gin.Context) {
	var recipes []models.Recipe

	a.db.Where("building_id = ?", g.Param("building_id")).Find(&recipes)

	g.JSON(http.StatusOK, gin.H{
		"recipes": recipes,
	})
}
