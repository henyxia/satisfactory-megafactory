package models

import (
	"fmt"
	"strings"
)

type ItemType string

const (
	ItemTypeItem   = "item"
	ItemTypeLiquid = "liquid"
	ItemTypeGas    = "gas"
)

type Item struct {
	ID   uint
	Name string
	Type ItemType
}

func (i *Item) URLImage() string {
	name := i.Name
	name = strings.ReplaceAll(name, " ", "-")
	name = strings.ToLower(name)
	return fmt.Sprintf("/assets/img/item-%s.png", name)
}
