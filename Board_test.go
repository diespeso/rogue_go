package rogue

import (
	"fmt"
	"testing"
)

type PlayerTest struct {
	Entity
}

type Enemy struct {
	Entity
}

var entities = map[int]EntityObject{}

func TestAll(t *testing.T) {
	newBoard := NewBoard(5)

	newBoard.FillWith(GROUND)
	player := newBoard.NewEntity(Position{0, 0}, PLAYER)
	entities[0] = player

	player.SetPosition(Position{0, 0})

	newBoard.update()
	newBoard.Show()

	player.MoveTo(Position{1, 0})
	player.Update()
	newBoard.update()
	newBoard.Show()

	player.MoveTo(Position{2, 0})
	player.Update()
	newBoard.update()
	newBoard.Show()

	player.MoveTo(Position{3, 0})
	player.Update()
	newBoard.update()
	newBoard.Show()

	player.board.GetCell(player.position.x, player.position.y).ShowInfo()

}




