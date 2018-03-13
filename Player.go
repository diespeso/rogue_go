package rogue

import "fmt"

type Health int

func (health Health) decrease(amount Health) {
	health -= amount
}

func (health Health) increase(amount Health) {
	health += amount
}

type PlayerState int
type Direction int

const (
	ALIVE PlayerState = 0
	DEAD  PlayerState = 1
)

const (
	UP    Direction = 0
	DOWN  Direction = 1
	LEFT  Direction = 2
	RIGHT Direction = 3
)

type Position struct {
	x int
	y int
}

type Player struct {
	Cell
	health   Health
	attack   Health
	state    PlayerState
	position Position
}

func (playerState PlayerState) CastToString() string {
	switch playerState {
	case ALIVE:
		return "alive"
	case DEAD:
		return "dead"
	}
	return "unknown"
}

func (player *Player) Show() {
	fmt.Printf("[Cell:%v, health:%v, attack:%v, state:%v, position(%v,%v)]",
		player.Cell, player.health, player.attack, player.state.CastToString(),
		player.position.x, player.position.y)
}

func NewPlayer() *Player {
	return &Player{NewCell(PLAYER), Health(100), Health(10), ALIVE, Position{0, 0}}
}

func (player *Player) Move(direction Direction) {
	switch direction {
	case UP:
		player.position.y -= 1
	case DOWN:
		player.position.y += 1
	case LEFT:
		player.position.x -= 1
	case RIGHT:
		player.position.x += 1
	}
}
