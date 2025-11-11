package ui

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/henyxia/satisfactory-megafactory/models"
)

func (u *UI) ItemList(g *gin.Context) {
	var items []models.Item

	u.db.Find(&items)

	g.HTML(http.StatusOK, "pages/item-list.html", gin.H{
		"items": items,
	})
}
