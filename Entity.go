package rogue

//An Entity is a object that exists on a Board and has a cellUnder it
type Entity struct {
	Cell
	position  Position
	health    Health
	attack    Health
	cellUnder CellObject
	board     *Board
}

//SetPosition sets the position of the Entity
func (entity *Entity) SetPosition(position Position) {
	entity.position = position
}

//GetPosition returns the position of the Entity
func (entity *Entity) GetPosition() Position {
	return entity.position
}

//ShowInfo shows debug info of the Entity
func (entity *Entity) ShowInfo() {
	//placeholder
	fmt.Printf("[Cell: ")
	entity.Cell.ShowInfo()

	fmt.Printf(", Position: %v, health: %v, attack:%v,",
		entity.position, entity.health, entity.attack)
	entity.cellUnder.ShowInfo()
	fmt.Print("]")
}

//Update updates the cellUnder the Entity
func (entity *Entity) Update() {
	entity.cellUnder = entity.board.GetCell(entity.position.x, entity.position.y)
	if entity.cellUnder.GetCellType() == GROUND {	//debug
		fmt.Println("Omg i'm stepping on dirt")	//debug
	} //debug
}

//MoveTo moves the Entity to a new Position and Updates the entity
//TODO: the Update part must be revised to see if is good to call from here
func (entity *Entity) MoveTo(position Position) {
	entity.board.GetCell(entity.position.x, entity.position.y).ShowInfo()
	entity.board.SetPCell(entity.position.x, entity.position.y, entity.cellUnder) //use SetPCell instead
	entity.position.x = position.x
	entity.position.y = position.y
	entity.Update()
}
