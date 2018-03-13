package rogue

//A CellType holds the type for a Cell
type CellType int

//A CellStruct holds the type information to be used on the cells map
type CellStruct struct {
	cellType  CellType
	name      string
	character string
}

//Enum of the CellType's
const (
	VOID   CellType = 0
	GROUND CellType = 1
	WATER  CellType = 2
	PLAYER CellType = 3
)

//cells holds the CellStructs to be called to access a CellType data
var cells map[CellType]CellStruct

//Initializes the map cells, without this there'd be a initialization loop
//https://stackoverflow.com/questions/31726254/how-to-avoid-initialization-loop-in-go
func init() {
	cells = map[CellType]CellStruct{
		VOID:   CellStruct{VOID, "void", " "},
		GROUND: CellStruct{GROUND, "ground", "-"},
		WATER:  CellStruct{WATER, "water", "*"},
		PLAYER: CellStruct{PLAYER, "player", "@"},
	}
}

//CastToString returns the string name of the cellType
func (cellType CellType) CastToString() string {
	return cells[cellType].name
}

//CastToCharacter returns the character representation of the cellType
func (cellType CellType) CastToCharacter() string {
	return cells[cellType].character
}

//CastToString casts a Celltype type to its string form(its name)
// func (cellType CellType) CastToString() string {
// 	switch cellType {
// 	case VOID:
// 		return "Void"
// 	case GROUND:
// 		return "Ground"
// 	case WATER:
// 		return "Water"
// 	case PLAYER:
// 		return "Player"
// 	default:
// 		return "Void"
// 	}
// }

//CastToCharacter casts a Celltype to its character representation
// func (cellType CellType) CastToCharacter() string {
// 	switch cellType {
// 	case VOID:
// 		return " "
// 	case GROUND:
// 		return "-"
// 	case WATER:
// 		return "+"
// 	case PLAYER:
// 		return "@"
// 	default:
// 		return "\\"
// 	}
// }
