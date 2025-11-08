package ui

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *UI) ErrNotFound(g *gin.Context) {
	g.HTML(http.StatusNotFound, "404.html", gin.H{})
}
