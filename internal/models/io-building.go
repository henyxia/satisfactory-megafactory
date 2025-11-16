package models

type IOBuilding struct {
	ID         uint
	Type       uint `min: 0 max:1`
	Building   Building
	BuildingID uint `json:"building_id"`
	Recipe     Recipe
	RecipeSlug string `json:"recipe_slug"`
	Note       string `json:"note"`
}

var ioBuildingTypeAsString = map[uint]string{
	0: "Producer",
	1: "Consumer",
}

func (b *IOBuilding) TypeAsString() string {
	return ioBuildingTypeAsString[b.Type]
}
