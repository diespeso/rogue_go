package rogue

import "fmt"

//A Cell type has a position x and y on a board, is shown ingame
type Cell struct {
	character string
	cellType  CellType
}

//NewCell returns a new cell initialized with the specified CellType
func NewCell(cellType CellType) Cell {
	//return Cell{cellType.CastToCharacter(), cellType}
	return Cell{cells[cellType].character, cellType}
}

//ShowInfo shows the information of the cell for debug purposes */
func (cell Cell) ShowInfo() {
	fmt.Printf("[%v, %v]", cell.character, cell.cellType.CastToString())
}

//SetCharacter sets the caracter representation for the Cell, this could not be used
func (cell Cell) SetCharacter(character string) {
	cell.character = character
}

//GetCharacter Returns the character that represents the Cell
func (cell Cell) GetCharacter() string {
	return cell.character
}

//SetCellType sets the CellType for the Cell */
func (cell Cell) SetCellType(cellType CellType) {
	cell.cellType = cellType
}

//GetCellType returns the CellType of the Cell */
func (cell Cell) GetCellType() CellType {
	return cell.cellType
}

//OnEnter is called when the player enters the cell
func (cell Cell) OnEnter() {
	fmt.Println("NO behaviour defined")
}

//CanEnter returns wheter the cell is able to be entered or not
func (cell Cell) CanEnter() bool {
	return false
}
