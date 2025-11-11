package models

type RecipeInput struct {
	RecipeID string `gorm:"primaryKey;uniqueIndex:idx_recipe_input_item"`
	Item     Item
	ItemID   uint `gorm:"primaryKey;uniqueIndex:idx_recipe_input_item"`
	Qt       uint
}

type RecipeOutput struct {
	RecipeID string `gorm:"primaryKey;uniqueIndex:idx_recipe_output_item"`
	Item     Item
	ItemID   uint `gorm:"primaryKey;uniqueIndex:idx_recipe_output_item"`
	Qt       uint
}

type Recipe struct {
	Slug       string `gorm:"primary_key"`
	Name       string
	Duration   uint     // duration as tenth of a second
	ProducedIn Building `gorm:"foreignKey:BuildingID"`
	BuildingID uint
	Input      []RecipeInput
	Output     []RecipeOutput
}
