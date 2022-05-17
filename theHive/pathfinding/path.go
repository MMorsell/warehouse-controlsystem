package pathfinding

import (
	"container/heap"
)

func heuristic(a Position, b Position) int32 {
	return abs(a.X-b.X) + abs(a.Y-b.Y)
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

		if current.X == goal.X && current.Y == goal.Y {

			// A bug may appear here with time, but I'm not sure.
			// Since we do not know time of arrival we must check time here too actually
			// but that would require us to know how long it would take to get here.
			// Does neighbour function take care of this already? I think it might.
			// Otherwise using current cost might solve this...somehow.
			// Conclusion: might be a bug here unsure, requires testing.

			// Add the time for how long it took so we know in the future when it will arrive.
			return constructPath(start, Position{goal.X, goal.Y, costSoFar[current]}, cameFrom)
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

func ReservePath(grid *Grid, path []Position) {
	for _, pos := range path {
		// We need to reserve the actual path but also the next time step to avoid head on collisions.
		// This is because two agents cannot actually interchange position but must traverse from one coordinate to another
		// and therefore may crash into each other. Think of this as also allowing for the time for an agent to leave its current position
		// before another agent can enter that coordinate.
		// More strict reasoning follows here otherwise:
		// Agent 1 reserves (x, y, t) and (x + 1, y, t + 1) which leaves another agent, say Agent 2, open to reserve (x + 1, y, t) and (x, y, t + 1)
		// which would result in a crash since they cannot travel through each other.
		// Visual example:
		// Agent 1: # wants to travel from (1, 1, 0) to (4, 1, 3) so it reserves (2, 1, 1), (3, 1, 2) and (4, 1, 3)
		// Agent 2: x wants to travel from (4, 1, 1) to (1, 1, 3) so it reserves (3, 1, 2), (2, 1, 3) and (1, 1, 4)
		// Look what happens when # is at (2, 1, 1) and x is at (3, 1, 2); there is a collision since the algorithm does not take into account
		// the actual traversal between coordinates but thinks any movement is like a teleport.
		// . . . . . .
		// . # # x x .
		// . . . . . .
		grid.Reserve(pos)
		grid.Reserve(Position{pos.X, pos.Y, pos.T + 1}) // Also reserve the time spot after.
	}
}

func reverse(s []Position) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
