package server

import (
	_ui "github.com/henyxia/satisfactory-megafactory/ui"
)

func (s *Server) setupRoutes() {
	ui := _ui.New(s.DB)

	// ui routes
	s.router.GET("/", ui.Index)
	s.router.GET("/floor/:id", ui.FloorView)
	s.router.GET("/floor/:id/plan", ui.FloorPlan)
	s.router.GET("/item", ui.ItemList)
	s.router.GET("/recipes", ui.RecipeList)
}
