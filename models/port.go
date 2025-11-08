package models

import "fmt"

type PortSubPosition string

const (
	PortSubPositionInnerLeft  = "inner-left"
	PortSubPositionInnerRight = "inner-right"
	PortSubPositionOuterLeft  = "outer-left"
	PortSubPositionOuterRight = "outer-right"
)

type Port struct {
	ID          uint
	FloorID     uint
	Position    uint
	SubPosition PortSubPosition
}

var subPositionAsString = map[PortSubPosition]string{
	PortSubPositionInnerLeft:  "Inner Left",
	PortSubPositionInnerRight: "Inner Right",
	PortSubPositionOuterLeft:  "Outer Left",
	PortSubPositionOuterRight: "Outer Right",
}

func (p *Port) SubPositionAsString() string {
	return subPositionAsString[p.SubPosition]
}

func (p *Port) String() string {
	return fmt.Sprintf("Level %d Position %d %s", p.FloorID, p.Position, p.SubPositionAsString())
}
