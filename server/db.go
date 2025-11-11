package server

import "github.com/henyxia/satisfactory-megafactory/models"

func (s *Server) migrateDB() error {
	s.DB.AutoMigrate(
		models.Factory{},
		models.Floor{},
		models.ProductionLine{},
		models.Port{},
		models.Building{},
		models.Item{},
		models.Recipe{},
		models.RecipeInput{},
		models.RecipeOutput{},
	)

	s.setData()

	return nil
}
