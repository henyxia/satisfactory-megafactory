package ui

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/henyxia/satisfactory-megafactory/models"
)

func (u *UI) RecipeList(g *gin.Context) {
	recipes := []models.Recipe{}

	u.db.Preload("Input.Item").Preload("Output.Item").Preload("ProducedIn").Find(&recipes)

	g.HTML(http.StatusOK, "pages/recipe-list.html", gin.H{
		"recipes": recipes,
	})
}
