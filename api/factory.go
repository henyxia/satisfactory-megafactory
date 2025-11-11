package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/henyxia/satisfactory-megafactory/models"
)

func (a *API) FactoryNew(g *gin.Context) {
	var factory models.Factory

	err := g.ShouldBindJSON(&factory)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = a.db.Save(&factory).Error
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusAccepted, gin.H{})
}
