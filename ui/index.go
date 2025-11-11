package ui

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/henyxia/satisfactory-megafactory/models"
)

func (u *UI) Index(g *gin.Context) {
	/*
		var floors []models.Floor

		u.db.Find(&floors)
	*/

	var factories []models.Factory

	u.db.Find(&factories)

	g.HTML(http.StatusOK, "pages/index.html", gin.H{
		"factories": factories,
	})
}
