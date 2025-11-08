package models

type ProductionLine struct {
	ID           uint
	Number       uint
	FloorID      uint
	PortInput    *Port
	PortInputID  uint
	PortOutput   *Port
	PortOutputID uint
}
