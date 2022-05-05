package pathfinding

import (
	"sort"
)

// Position struct defines a position on the grid in the form of (x, y) coordinates.
type Position struct {
	x, y int32
}

type Grid struct {
	width      int32
	height     int32
	obstructed map[Position]struct{}
	reserved   map[Position][]uint32
}

func NewGrid(width int32, height int32, obstacles []Position) *Grid {
	obstacleMap := make(map[Position]struct{}, len(obstacles))
	reservedMap := make(map[Position][]uint32)

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

// Traversable returns all whether the rovided Position is traversable at provided time.
// A position is traversable if it is contained in the grid, not obstructed and not reserved.
func (g *Grid) Traversable(pos Position, time uint32) bool {
	_, exists := g.obstructed[pos]
	return g.Within(pos) && !exists && !g.Reserved(pos, time)
}

// Reserved returns true if provided Position is reserved at provided time.
// This should only ever be used by an agent who needs to reserve this Position for their path
// Note that non-reserved Positions may still contain obstacles and return false.
func (g *Grid) Reserved(pos Position, time uint32) bool {
	arr := g.reserved[pos]
	index := sort.Search(len(arr), func(i int) bool { return time <= arr[i] })
	return index < len(arr) && arr[index] == time
}

// Reserve markes the provided Position as reserved at provided time.
func (g *Grid) Reserve(pos Position, time uint32) {
	arr := g.reserved[pos]
	index := sort.Search(len(arr), func(i int) bool { return time <= arr[i] })
	g.reserved[pos] = insertAt(index, time, arr)
}

// TraversableNeighbours returns all neighbours to provided Position that are traversable at provided time.
// A neighbour is traversable if it is contained in the grid, not obstructed and not reserved.
func (g *Grid) TraversableNeighbours(pos Position, time uint32) []Position {

	var neighbours []Position

	for _, pos := range []Position{{pos.x + 1, pos.y}, {pos.x - 1, pos.y}, {pos.x, pos.y - 1}, {pos.x, pos.y + 1}} {
		if g.Traversable(pos, time) {
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
