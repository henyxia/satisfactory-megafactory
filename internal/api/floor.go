package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/henyxia/satisfactory-megafactory/internal/models"
)

func (a *API) FloorNew(g *gin.Context) {
	var floor models.FloorNew

	err := g.ShouldBindJSON(&floor)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = a.db.Create(&floor).Error
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusAccepted, gin.H{})
}
