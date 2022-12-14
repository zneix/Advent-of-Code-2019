package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// part 1 = 2
// part 2 = 10
const KNOT_COUNT int = 10

var (
	knots         = make([]coordinate, KNOT_COUNT, KNOT_COUNT)
	visitedByTail = make(map[coordinate]bool)

	// used for prints
	minY, maxY, minX, maxX = 0, 0, 0, 0
)

type coordinate struct {
	x, y int
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func signum(x int) int {
	switch {
	case x > 0:
		return 1
	case x < 0:
		return -1
	default:
		return 0
	}
}

func moveHead(direction byte, head *coordinate) {
	switch direction {
	case 'R':
		head.x++
	case 'L':
		head.x--
	case 'U':
		head.y++
	case 'D':
		head.y--
	}

	// used for prints
	if head.x > maxX {
		maxX = head.x
	}
	if head.x < minX {
		minX = head.x
	}
	if head.y > maxY {
		maxY = head.y
	}
	if head.y < minY {
		minY = head.y
	}
}

func moveTail(head, tail *coordinate) {
	deltaX := head.x - tail.x
	deltaY := head.y - tail.y

	if abs(deltaX) > 1 || abs(deltaY) > 1 {
		tail.x += signum(deltaX)
		tail.y += signum(deltaY)
	}
}

func main() {
	input, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		direction := line[0]
		amount, _ := strconv.Atoi(string(line[2:]))

		// move head N times in specified direction and then all of its tails
		for i := 0; i < amount; i++ {
			moveHead(direction, &knots[0])
			for j := 0; j < KNOT_COUNT-1; j++ {
				moveTail(&knots[j], &knots[j+1])
			}
			visitedByTail[knots[KNOT_COUNT-1]] = true
		}
	}
	fmt.Println(len(visitedByTail))

	output, err := os.Create("output")
	if err != nil {
		panic(err)
	}
	for y := maxY; y >= minY; y-- {
		line := make([]byte, 0)
		for x := minX; x <= maxX; x++ {
			switch {
			case knots[0].y == y && knots[0].x == x:
				line = append(line, 'H')
			case knots[KNOT_COUNT-1].y == y && knots[KNOT_COUNT-1].x == x:
				line = append(line, '9')
			//case x == 0 && y == 0:
			//line += "s"
			case visitedByTail[coordinate{y: y, x: x}]:
				line = append(line, '#')
			default:
				line = append(line, '.')
			}
		}
		output.Write(append(line, '\n'))
	}
}
