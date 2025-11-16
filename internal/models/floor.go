package models

import "fmt"

type FloorStatus string

const (
	FloorStatusAbsent     FloorStatus = "absent"
	FloorStatusInProgress FloorStatus = "in-progress"
	FloorStatusDone       FloorStatus = "done"
)

type FloorType string

const (
	FloorTypeAbsent      FloorType = "absent"
	FloorTypeSmelter     FloorType = "smelter"
	FloorTypeConstructor FloorType = "constructor"
	FloorTypeAssembler   FloorType = "assembler"
)

type Floor struct {
	ID              uint `uri:"floor_id" binding:"required"`
	Factory         Factory
	FactoryID       uint        `json:"factory_id" gorm:"uniqueIndex:idx_floor_factory_level"`
	Level           int         `json:"level" gorm:"uniqueIndex:idx_floor_factory_level"`
	Name            string      `json:"name"`
	Status          FloorStatus `json:"status"`
	Type            FloorType
	ProductionLines []ProductionLine
	Ports           []Port
	IOs             []IO
}

var floorTypeAsUIClass = map[FloorType]string{
	FloorTypeAbsent:      "",
	FloorTypeSmelter:     "is-success is-light",
	FloorTypeConstructor: "is-info is-light",
	FloorTypeAssembler:   "is-warning is-light",
}

func (f *Floor) TypeUIClass() string {
	return floorTypeAsUIClass[f.Type]
}

var floorStatusAsUIClass = map[FloorStatus]string{
	FloorStatusAbsent:     "is-light",
	FloorStatusInProgress: "is-info",
	FloorStatusDone:       "is-success",
}

func (f *Floor) StatusUIClass() string {
	return floorStatusAsUIClass[f.Status]
}

func (f *Floor) URLView() string {
	return fmt.Sprintf("/factory/%d/floor/%d", f.FactoryID, f.ID)
}

func (f *Floor) BreadcrumbName() string {
	return fmt.Sprintf("%d - %s", f.Level, f.Name)
}

type FloorNew struct {
	FactoryID uint        `json:"factory_id" binding:"required" gorm:"uniqueIndex:idx_floor_factory_level"`
	Level     int         `json:"level" binding:"required" gorm:"uniqueIndex:idx_floor_factory_level"`
	Name      string      `json:"name" binding:"required"`
	Status    FloorStatus `json:"status" binding:"required"`
}

func (FloorNew) TableName() string {
	return "floors"
}
