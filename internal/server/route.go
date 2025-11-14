package server

import (
	_api "github.com/henyxia/satisfactory-megafactory/internal/api"
	_ui "github.com/henyxia/satisfactory-megafactory/internal/ui"
)

func (s *Server) setupRoutes() {
	api := _api.New(s.DB)

	// api routes
	s.router.POST("/api/factory", api.FactoryNew)
	s.router.DELETE("/api/factory/:factory_id", api.FactoryDelete)
	s.router.POST("/api/factory/floor", api.FloorNew)
	s.router.PUT("/api/factory/:factory_id/floor/:floor_id", api.FloorEdit)
	s.router.DELETE("/api/factory/:factory_id/floor/:floor_id", api.FloorDelete)

	ui := _ui.New(s.DB)

	// ui routes
	s.router.GET("/", ui.FactoryList)
	s.router.GET("/factory/:factory_id", ui.FactoryView)
	s.router.GET("/factory/:factory_id/floor/:floor_id", ui.FloorView)
	s.router.GET("/floor/:floor_id/plan", ui.FloorPlan)
	s.router.GET("/item", ui.ItemList)
	s.router.GET("/recipes", ui.RecipeList)
}
