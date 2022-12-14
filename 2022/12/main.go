package main

import (
	"fmt"
	"os"
	"strings"
)

type coord struct {
	x, y int
}
type tileQueue []coord

// start has elevation 'a', end has elevation 'z'
var (
	grid         = make(map[coord]int) // int: altitude/height of coord which is letters 'a' - 'z'
	end, start   coord
	MAX_X, MAX_Y int
)

func (self *tileQueue) tryPush(visited map[coord]int, current coord, next coord, height int) {
	// going forward, can climb 1, fall many (grid[next] > height+1)
	// going backwards, can climb many, fall 1
	if grid[next]+1 < height {
		return
	}

	_, has := visited[next] // don't care about the value
	if !has {
		*self = append(*self, next)
		//fmt.Printf("setting %v to value of %v (%d+1)\n", next, current, visited[current])
		visited[next] = visited[current] + 1
	}
}

func findSteps(first coord) map[coord]int {
	//fmt.Printf("start=%v, end=%v\n", start, end)
	visited := make(map[coord]int) // int: distance from 'S' to coord
	visited[first] = 0

	curr := first
	queue := tileQueue{curr}
	for len(queue) > 0 {
		//fmt.Println("tiles left", len(queue))
		// shift the queue slice
		curr = queue[0]
		//if curr == end {
		////fmt.Println("we found end", end, visited[end])
		////break
		//}
		queue = queue[1:]

		// get altitude of the tile we're examining
		// to know if we can step on the tile next to it
		currHeight := grid[curr]

		// try to move up
		nextY := curr.y + 1
		if nextY >= 0 && nextY <= MAX_Y {
			next := coord{x: curr.x, y: nextY}
			queue.tryPush(visited, curr, next, currHeight)
		}
		// try to move down
		nextY = curr.y - 1
		if nextY >= 0 && nextY <= MAX_Y {
			next := coord{x: curr.x, y: nextY}
			queue.tryPush(visited, curr, next, currHeight)
		}
		// try to move right
		nextX := curr.x + 1
		if nextX >= 0 && nextX <= MAX_X {
			next := coord{x: nextX, y: curr.y}
			queue.tryPush(visited, curr, next, currHeight)
		}
		// try to move left
		nextX = curr.x - 1
		if nextX >= 0 && nextX <= MAX_X {
			next := coord{x: nextX, y: curr.y}
			queue.tryPush(visited, curr, next, currHeight)
		}
	}

	//fmt.Println("idk???", MAX_X, MAX_Y, (MAX_X+1)*(MAX_Y+1), len(visited), visited[end])
	return visited
}

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	// bottom-left character will be coord{0, 0}
	lines := strings.Fields(string(data))
	MAX_Y = len(lines) - 1
	MAX_X = len(lines[0]) - 1

	for y, i := 0, MAX_Y; i >= 0; y, i = y+1, i-1 {
		for x, c := range lines[i] {
			elevation := int(c)
			switch c {
			case 'S':
				elevation = 'a'
				start = coord{x: x, y: y}
			case 'E':
				elevation = 'z'
				end = coord{x: x, y: y}
			}
			grid[coord{x: x, y: y}] = elevation
		}
	}

	visitedSteps := findSteps(end)
	fmt.Println(visitedSteps[start])
	leastSteps := 1337 // random high number

	// XXX: hacky solution
	for y := MAX_Y; y >= 0; y-- {
		steps := visitedSteps[coord{x: start.x, y: y}]
		if steps < leastSteps {
			leastSteps = steps
		}
	}
	fmt.Println(leastSteps)
}
