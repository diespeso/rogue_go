//Package rogue holds the pieces to build a roguelike game
package rogue

import (
	"fmt"
)

//A Board holds the space to play
type Board struct {
	size   int
	matrix [][]CellObject
}

//NewBoard returns a new Board with its matrix initialized */
func NewBoard(size int) Board {
	return Board{size, GetNewMatrix(size)}
}

//GetNewMatrix returns a CellObject matrix
func GetNewMatrix(size int) [][]CellObject {
	newMatrix := make([][]CellObject, size)
	for i := range newMatrix {
		newMatrix[i] = make([]CellObject, size)
		for e := range newMatrix[i] {
			newMatrix[i][e] = NewVoid()
		}
	}
	return newMatrix
}

//Show shows the characters of the board objects (their sprites)
func (board *Board) Show() {
	for i := range board.matrix {
		for e := range board.matrix[i] {
			fmt.Printf("%v ", board.matrix[i][e].GetCharacter())
		}
		fmt.Println()
	}
}

//FillWith fills the board with the specified cell
func (board *Board) FillWith(cellType CellType) {
	for i := range board.matrix {
		for e := range board.matrix[i] {
			board.matrix[i][e] = &Cell{cellType.CastToCharacter(), cellType}

		}
	}
}

//SetCell sets the board[x][y] to the a cell of the cellType
func (board *Board) SetCell(x int, y int, cellType CellType) {
	board.matrix[x][y] = NewCell(cellType)
}

func (board *Board) SetPCell(x int, y int, cell CellObject) {
	board.matrix[x][y] = cell
}

func (board *Board) SetCellObject(x int, y int, cellObject CellObject) {
	board.matrix[x][y] = cellObject
}

//GetCell return the board[x][y] cellObject
func (board *Board) GetCell(x int, y int) CellObject {
	return board.matrix[x][y]
}
