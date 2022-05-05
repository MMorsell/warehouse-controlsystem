package pathfinding

import (
	"container/heap"
)

func heuristic(a Position, b Position) int32 {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func abs(x int32) int32 {
	m := x >> 31
	return (x ^ m) - m
}

func FindPath(grid *Grid, start Position, goal Position) []Position {

	costSoFar := make(map[Position]int32)
	cameFrom := make(map[Position]Position)

	costSoFar[start] = 0
	cameFrom[start] = start

	q := &priorityQueue{}
	heap.Init(q)
	heap.Push(q, &item{value: start, priority: 0})

	for q.Len() > 0 {
		current := heap.Pop(q).(*item).value

		if current.x == goal.x && current.y == goal.y {
			return constructPath(start, Position{goal.x, goal.y, costSoFar[current]}, cameFrom)
		}

		for _, neighbour := range grid.TraversableNeighbours(current) {
			newCost := costSoFar[current] + 1
			cost, exists := costSoFar[neighbour]
			if !exists || (newCost < cost) {
				costSoFar[neighbour] = newCost
				cameFrom[neighbour] = current
				heap.Push(q, &item{value: neighbour, priority: newCost + heuristic(neighbour, goal)})
			}
		}
	}
	return nil
}

func constructPath(start Position, goal Position, cameFrom map[Position]Position) []Position {
	var path []Position
	current := goal
	for current != start {
		path = append(path, current)
		current = cameFrom[current]
	}
	reverse(path)
	return path
}

func reverse(s []Position) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
