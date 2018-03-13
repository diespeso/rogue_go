package rogue

import (
	"fmt"
	"testing"
)

type Entity struct {
	Cell
	position  Position
	health    Health
	attack    Health
	cellUnder CellObject
	board     *Board
}

type EntityObject interface {
	CellObject
	SetPosition(position Position)
	GetPosition() Position
}

func (entity *Entity) SetPosition(position Position) {
	entity.position = position
}

func (entity *Entity) GetPosition() Position {
	return entity.position
}

func (entity *Entity) ShowInfo() {
	//placeholder
	fmt.Printf("[Cell: ")
	entity.Cell.ShowInfo()

	fmt.Printf(", Position: %v, health: %v, attack:%v,",
		entity.position, entity.health, entity.attack)
	entity.cellUnder.ShowInfo()
	fmt.Print("]")
}

func (board *Board) NewEntity(position Position, cellType CellType) *Entity {
	return &Entity{
		Cell{cells[cellType].character, cellType},
		Position{position.x, position.y},
		Health(100),
		Health(10),
		board.GetCell(position.x, position.y),
		board,
	}
}

func (entity *Entity) Update() {
	entity.cellUnder = entity.board.GetCell(entity.position.x, entity.position.y)
	if entity.cellUnder.GetCellType() == GROUND {
		fmt.Println("Omg i'm stepping on dirt")
	}
}

// func (entity Entity) SetCharacter(character string) {
// 	entity.character = character
// }

// func (entity Entity) GetCharacter() string {
// 	return entity.character
// }

// func (entity Entity) SetCellType(cellType CellType) {
// 	entity.cellType = cellType
// }

// func (entity Entity) GetCellType() CellType {
// 	return entity.cellType
// }

// func (entity Entity) OnEnter() {
// 	//placeholder but should throw an error
// }

// func (entity Entity) CanEnter() bool {
// 	return false
// }

// // type EntityObject interface {
// // 	CellObject
// // }

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

func (entity *Entity) MoveTo(position Position) {
	entity.board.GetCell(entity.position.x, entity.position.y).ShowInfo()
	entity.board.SetPCell(entity.position.x, entity.position.y, entity.cellUnder) //use SetPCell instead
	entity.position.x = position.x
	entity.position.y = position.y
	entity.Update()

}

func (board *Board) update() {
	for _, key := range entities {
		board.SetCellObject(key.GetPosition().x, key.GetPosition().y, key)
		fmt.Println(key.GetPosition().x, key.GetPosition().y)
	}
}
