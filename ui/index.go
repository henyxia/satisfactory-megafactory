package ui

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/henyxia/satisfactory-megafactory/models"
)

func (u *UI) Index(g *gin.Context) {
	var floors []models.Floor

	u.db.Find(&floors)

	g.HTML(http.StatusOK, "index.html", gin.H{
		"floors": floors,
	})
}
