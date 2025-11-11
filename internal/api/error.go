package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *API) ErrNotFound(g *gin.Context) {
	g.JSON(http.StatusNotFound, gin.H{"error": "id might be incorrect"})
}
