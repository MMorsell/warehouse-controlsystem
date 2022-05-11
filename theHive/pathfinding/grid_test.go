package pathfinding

import (
	"fmt"
	"testing"
)

func initGridNoObstructions() *Grid {
	var obstructions []Position
	return NewGrid(10, 5, obstructions)
}

func TestWithin(t *testing.T) {

	grid := initGridNoObstructions()

	cases := []struct {
		pos  Position
		want bool
	}{
		{Position{5, 2, 0}, true},
		{Position{10, 5, 0}, false},
		{Position{-2, 3, 0}, false},
		{Position{5, 8, 0}, false},
		{Position{2, 2, 15}, true},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("(x=%d,y=%d,t=%d) with (width=%d,height=%d)", tc.pos.x, tc.pos.y, tc.pos.t, grid.width, grid.height), func(t *testing.T) {
			got := grid.Within(tc.pos)
			if tc.want != got {
				t.Errorf("Expected '%t', but got '%t'", tc.want, got)
			}
		})
	}
}

func TestNeighboursNoObstructions(t *testing.T) {
	grid := initGridNoObstructions()

	cases := []struct {
		pos  Position
		want []Position
	}{
		{Position{0, 0, 0}, []Position{{0, 1, 1}, {1, 0, 1}, {0, 0, 1}}},
		{Position{1, 0, 0}, []Position{{0, 0, 1}, {2, 0, 1}, {1, 1, 1}, {1, 0, 1}}},
		{Position{1, 1, 0}, []Position{{1, 0, 1}, {0, 1, 1}, {1, 2, 1}, {2, 1, 1}, {1, 1, 1}}},
		{Position{9, 4, 0}, []Position{{9, 3, 1}, {8, 4, 1}, {9, 4, 1}}},
		{Position{9, 4, 4}, []Position{{9, 3, 5}, {8, 4, 5}, {9, 4, 5}}},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("Neighbours for (x=%d,y=%d,t=%d)", tc.pos.x, tc.pos.y, tc.pos.t), func(t *testing.T) {
			got := grid.TraversableNeighbours(tc.pos)
			if !sliceHasSameElements(got, tc.want) {
				t.Errorf("Expected %v, but got %v", tc.want, got)
			}
		})
	}
}

func TestReservePosition(t *testing.T) {
	grid := initGridNoObstructions()

	cases := []struct {
		pos  []Position
		want []bool
	}{
		{[]Position{{5, 4, 1}, {5, 4, 4}}, []bool{true, true}},
		{[]Position{{1, 1, 1}}, []bool{true}},
		{[]Position{{2, 1, 5}, {2, 1, 2}}, []bool{true, true}},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("Reserve %v", tc.pos), func(t *testing.T) {
			for _, p := range tc.pos {
				grid.Reserve(p)
			}
			for i, p := range tc.pos {
				got := grid.Reserved(p)
				if got != tc.want[i] {
					t.Errorf("Expected %v, but got %v", tc.want, got)
				}
			}
		})
	}
}

func TestReservedPosition(t *testing.T) {
	grid := initGridNoObstructions()

	grid.Reserve(Position{1, 1, 4})
	grid.Reserve(Position{1, 1, 10})
	grid.Reserve(Position{5, 5, 4})

	cases := []struct {
		pos  Position
		want bool
	}{
		{Position{1, 1, 5}, false},
		{Position{1, 1, 4}, true},
		{Position{5, 5, 4}, true},
		{Position{1, 1, 1}, false},
		{Position{5, 5, 5}, false},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("Reserved %d", tc.pos), func(t *testing.T) {
			got := grid.Reserved(tc.pos)
			if got != tc.want {
				t.Errorf("Expected '%t', but got '%t'", tc.want, got)
			}
		})
	}
}

func sliceHasSameElements(x, y []Position) bool {
	if len(x) != len(y) {
		return false
	}
	diff := make(map[Position]int, len(x))

	for _, x_elem := range x {
		diff[x_elem]++
	}
	for _, y_elem := range y {
		if _, exists := diff[y_elem]; !exists {
			return false
		}
		diff[y_elem]--
		if diff[y_elem] == 0 {
			delete(diff, y_elem)
		}
	}
	return len(diff) == 0
}
