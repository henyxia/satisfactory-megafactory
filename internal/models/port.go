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
	ID          uint            `json:"id"`
	FloorID     uint            `json:"floor_id"`
	Direction   uint            `json:"direction"`
	Position    uint            `json:"position"`
	SubPosition PortSubPosition `json:"sub_position"`
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

var portDirectionAsString = map[uint]string{
	0: "North",
	1: "East",
	2: "South",
	3: "West",
}

func (p *Port) DirectionAsString() string {
	return ioDirectionAsString[p.Direction]
}
