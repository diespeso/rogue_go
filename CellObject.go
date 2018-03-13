//A Cell Object is a interface that holds the functionaltiy for a BoardObject

package rogue

//A CellObject holds the functionality for a Cell
type CellObject interface {
	SetCharacter(character string)
	GetCharacter() string

	SetCellType(cellType CellType)
	GetCellType() CellType

	ShowInfo()

	OnEnter()
	CanEnter() bool
}
