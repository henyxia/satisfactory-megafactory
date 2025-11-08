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
	ID              uint `uri:"id" binding:"required"`
	Name            string
	Status          FloorStatus
	Type            FloorType
	ProductionLines []ProductionLine
	Ports           []Port
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
	return fmt.Sprintf("/floor/%d", f.ID)
}

func (f *Floor) BreadcrumbName() string {
	return fmt.Sprintf("%d - %s", f.ID, f.Name)
}
