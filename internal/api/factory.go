package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/henyxia/satisfactory-megafactory/internal/models"
)

func (a *API) FactoryNew(g *gin.Context) {
	var factory models.FactoryNew

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

func (a *API) FactoryDelete(g *gin.Context) {
	factory := &models.Factory{}

	err := g.ShouldBindUri(factory)
	if err != nil {
		a.ErrNotFound(g)
		return
	}

	result := a.db.First(factory)

	if result.RowsAffected < 1 {
		a.ErrNotFound(g)
		return
	}

	err = a.db.Delete(&factory).Error
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusAccepted, gin.H{})
}
