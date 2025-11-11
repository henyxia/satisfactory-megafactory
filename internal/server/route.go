package server

import (
	_api "github.com/henyxia/satisfactory-megafactory/internal/api"
	_ui "github.com/henyxia/satisfactory-megafactory/internal/ui"
)

func (s *Server) setupRoutes() {
	api := _api.New(s.DB)

	// api routes
	s.router.POST("/api/factory", api.FactoryNew)
	s.router.DELETE("/api/factory/:id", api.FactoryDelete)

	ui := _ui.New(s.DB)

	// ui routes
	s.router.GET("/", ui.FactoryList)
	s.router.GET("/factory/:id", ui.FactoryView)
	s.router.GET("/floor/:id", ui.FloorView)
	s.router.GET("/floor/:id/plan", ui.FloorPlan)
	s.router.GET("/item", ui.ItemList)
	s.router.GET("/recipes", ui.RecipeList)
}
