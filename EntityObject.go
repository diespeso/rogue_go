package rogue

type EntityObject interface {
	CellObject //inherits from a cell object
	SetPosition(position Position)
	GetPosition() Position
}

