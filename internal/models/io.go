package models

type IO struct {
	ID          uint
	FloorID     uint       `json:"floor_id"`
	Direction   uint       `json:"direction"`
	Position    uint       `json:"position"`
	SubPosition uint       `json:"sub_position" min:0 max:3`
	BuildingID  uint       `json:"-"`
	Building    IOBuilding `gorm:"foreignKey:ID"`
}

var ioDirectionAsString = map[uint]string{
	0: "North",
	1: "East",
	2: "South",
	3: "West",
}

func (io *IO) DirectionAsString() string {
	return ioDirectionAsString[io.Direction]
}

var ioSubPositionAsString = map[uint]string{
	0: "Lower Right",
	1: "Higher Right",
	2: "Lower Left",
	3: "Higher Left",
}

func (io *IO) SubPositionAsString() string {
	return ioSubPositionAsString[io.SubPosition]
}
