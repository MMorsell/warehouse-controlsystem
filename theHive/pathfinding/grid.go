package pathfinding

// Position struct defines a space-time position on the grid in the form of (x, y, t) coordinates.
type Position struct {
	x, y, t int32
}

type Grid struct {
	width      int32
	height     int32
	obstructed map[Position]struct{}
	reserved   map[Position]struct{}
}

func NewGrid(width int32, height int32, obstacles []Position) *Grid {
	obstacleMap := make(map[Position]struct{}, len(obstacles))
	reservedMap := make(map[Position]struct{})

	for _, pos := range obstacles {
		obstacleMap[pos] = struct{}{}
	}

	return &Grid{width: width, height: height, obstructed: obstacleMap, reserved: reservedMap}
}

// Within returns true if provided Position is contained in the grid
// Note that positions with obstructions are contained in the grid and return true too.
func (g *Grid) Within(pos Position) bool {
	return pos.x >= 0 && pos.x < g.width && pos.y >= 0 && pos.y < g.height
}

// Obstructed returns true if provided Position is obstructed which is a time-independent property.
func (g *Grid) Obstructed(pos Position) bool {
	_, exists := g.obstructed[pos]
	return exists
}

// Reserved returns true if provided Position is reserved.
// This should only be used by an agent who needs to reserve this Position for their path
// Note that non-reserved Positions may still contain obstacles and return false.
func (g *Grid) Reserved(pos Position) bool {
	_, exists := g.reserved[pos]
	return exists
}

// Traversable returns true if the provided Position is traversable.
// A position is traversable if its space coordinates are contained in the grid, not obstructed and not reserved.
func (g *Grid) Traversable(pos Position) bool {
	return g.Within(pos) && !g.Obstructed(pos) && !g.Reserved(pos)
}

// Reserve markes the provided Position as reserved.
func (g *Grid) Reserve(pos Position) {
	g.reserved[pos] = struct{}{}
}

// TraversableNeighbours returns all neighbours to provided Position that are traversable.
// A neighbour is traversable if it is contained in the grid, not obstructed and not reserved.
func (g *Grid) TraversableNeighbours(pos Position) []Position {

	var neighbours []Position

	for _, pos := range []Position{right(pos), left(pos), up(pos), down(pos), wait(pos)} {
		if g.Traversable(pos) {
			neighbours = append(neighbours, pos)
		}
	}
	return neighbours
}

func insertAt(index int, value uint32, arr []uint32) []uint32 {
	if index == len(arr) {
		arr = append(arr, value)
	} else {
		// Retain sorted array by duplicating an element and then overwriting
		arr = append(arr[:index+1], arr[index:]...)
		arr[index] = value
	}
	return arr
}

func right(pos Position) Position {
	return Position{pos.x + 1, pos.y, pos.t + 1}
}

func left(pos Position) Position {
	return Position{pos.x - 1, pos.y, pos.t + 1}
}

func up(pos Position) Position {
	return Position{pos.x, pos.y - 1, pos.t + 1}
}

func down(pos Position) Position {
	return Position{pos.x, pos.y + 1, pos.t + 1}
}

func wait(pos Position) Position {
	return Position{pos.x, pos.y, pos.t + 1}
}
