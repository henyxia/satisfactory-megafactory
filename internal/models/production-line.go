package models

type ProductionLine struct {
	ID           uint  `json:"id"`
	Number       uint  `json:"number"`
	FloorID      uint  `json:"floor_id"`
	PortInput    *Port `json:"input"`
	PortInputID  uint
	PortOutput   *Port `json:"output"`
	PortOutputID uint
}
