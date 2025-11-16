package ui

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/henyxia/satisfactory-megafactory/internal/models"
)

func (u *UI) FloorView(g *gin.Context) {
	floor := &models.Floor{}

	err := g.ShouldBindUri(floor)
	if err != nil {
		u.ErrNotFound(g)
		return
	}

	result := u.db.
		Preload("Factory").
		Preload("ProductionLines").
		Preload("ProductionLines.PortInput").
		First(floor)

	if result.RowsAffected < 1 {
		u.ErrNotFound(g)
		return
	}

	g.HTML(http.StatusOK, "pages/floor-view.html", gin.H{
		"floor": floor,
	})
}

func (u *UI) FloorPlan(g *gin.Context) {
	floor := &models.Floor{}

	err := g.ShouldBindUri(floor)
	if err != nil {
		u.ErrNotFound(g)
		return
	}

	result := u.db.
		Preload("ProductionLines.PortInput").
		First(floor)

	if result.RowsAffected < 1 {
		u.ErrNotFound(g)
		return
	}

	megafactoryConfiguration := struct {
		NorthIO int
		EastIO  int
		SouthIO int
		WestIO  int
	}{
		NorthIO: 36,
		EastIO:  36,
		SouthIO: 36,
		WestIO:  36,
	}

	g.HTML(http.StatusOK, "pages/floor-plan.svg", gin.H{
		"floor": floor,
		"cfg":   megafactoryConfiguration,
	})
}

func (u *UI) FloorIO(g *gin.Context) {
	floor := &models.Floor{}

	err := g.ShouldBindUri(floor)
	if err != nil {
		u.ErrNotFound(g)
		return
	}

	result := u.db.
		Preload("ProductionLines.PortInput").
		First(floor)

	if result.RowsAffected < 1 {
		u.ErrNotFound(g)
		return
	}

	megafactoryConfiguration := struct {
		NorthIO int
		EastIO  int
		SouthIO int
		WestIO  int
	}{
		NorthIO: 36,
		EastIO:  36,
		SouthIO: 36,
		WestIO:  36,
	}

	g.HTML(http.StatusOK, "pages/floor-io.svg", gin.H{
		"floor": floor,
		"cfg":   megafactoryConfiguration,
	})
}
