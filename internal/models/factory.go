package models

type Factory struct {
	ID      uint
	Name    string `json:"name"`
	IONorth uint   `json:"io_north"`
	IOEast  uint   `json:"io_east"`
	IOSouth uint   `json:"io_south"`
	IOWest  uint   `json:"io_west"`
	Floors  []Floor
}
