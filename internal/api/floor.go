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

func (a *API) FloorEdit(g *gin.Context) {
	var floor models.Floor

	err := g.ShouldBindJSON(&floor)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = a.db.Save(&floor).Error
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusAccepted, gin.H{})
}

func (a *API) FloorDelete(g *gin.Context) {
	floor := &models.Floor{}

	err := g.ShouldBindUri(floor)
	if err != nil {
		a.ErrNotFound(g)
		return
	}

	result := a.db.First(floor)

	if result.RowsAffected < 1 {
		a.ErrNotFound(g)
		return
	}

	err = a.db.Delete(&floor).Error
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusAccepted, gin.H{})
}

func (a *API) FloorIO(g *gin.Context) {
	var ios []models.IO

	a.db.Where("floor_id = ?", g.Param("floor_id")).
		Preload("Building").
		Preload("Building.Recipe").
		Preload("Building.Recipe.Output").
		Preload("Building.Building").
		Find(&ios)

	g.JSON(http.StatusOK, gin.H{
		"ios": ios,
	})
}

func (a *API) FloorIONew(g *gin.Context) {
	input := struct {
		IODirection        uint   `json:"direction"`
		IOPosition         uint   `json:"position"`
		IOSubPosition      uint   `json:"sub_position"`
		BuildingID         uint   `json:"building_id" binding:"required"`
		BuildingRecipeSlug string `json:"recipe_slug" binding:"required"`
		BuildingNote       string `json:"note"`
	}{}

	err := g.ShouldBindJSON(&input)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error":      err.Error(),
			"additional": "unable to parse json",
		})
		return
	}

	floor := &models.Floor{}

	err = g.ShouldBindUri(floor)
	if err != nil {
		a.ErrNotFound(g)
		return
	}

	result := a.db.First(floor)

	if result.RowsAffected < 1 {
		a.ErrNotFound(g)
		return
	}

	building := &models.IOBuilding{
		BuildingID: input.BuildingID,
		RecipeSlug: input.BuildingRecipeSlug,
		Note:       input.BuildingNote,
	}

	io := &models.IO{
		FloorID:     floor.ID,
		Direction:   input.IODirection,
		Position:    input.IOPosition,
		SubPosition: input.IOSubPosition,
	}

	tx := a.db.Begin()
	err = tx.Create(building).Error
	if err != nil {
		tx.Rollback()
		g.JSON(http.StatusBadRequest, gin.H{
			"error":      err.Error(),
			"additional": "unable to register building in db",
		})
		return
	}

	io.Building = *building
	err = tx.Create(io).Error
	if err != nil {
		tx.Rollback()
		g.JSON(http.StatusBadRequest, gin.H{
			"error":      err.Error(),
			"additional": "unable to register io in db",
		})
		return
	}

	tx.Commit()

	g.JSON(http.StatusAccepted, gin.H{})
}
