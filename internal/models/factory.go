package models

import "fmt"

type Factory struct {
	ID      uint   `json:"factory_id" uri:"factory_id"`
	Name    string `json:"name"`
	IONorth uint   `json:"io_north"`
	IOEast  uint   `json:"io_east"`
	IOSouth uint   `json:"io_south"`
	IOWest  uint   `json:"io_west"`
	Floors  []Floor
}

func (f *Factory) URLView() string {
	return fmt.Sprintf("/factory/%d", f.ID)
}

func (f *Factory) String() string {
	return fmt.Sprintf("factory id:%d named: %s", f.ID, f.Name)
}

type FactoryNew struct {
	ID      uint   `uri:"factory_id"`
	Name    string `json:"name" binding:"required"`
	IONorth uint   `json:"io_north"`
	IOEast  uint   `json:"io_east"`
	IOSouth uint   `json:"io_south"`
	IOWest  uint   `json:"io_west"`
}

func (FactoryNew) TableName() string {
	return "factories"
}
