package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/// types

// WARNING: -x is left, +x is right BUT -y is up, +y is down
type coord struct {
	x, y int
}

func (self *coord) Add(other coord) {
	self.x += other.x
	self.y += other.y
}

type item rune
type grid struct {
	items                  map[coord]item
	MinX, MinY, MaxX, MaxY int
	FloorY                 int // defined just for the more readable code
	SandCount              int
	SandCountPart1         int // snapshotted once we reach a certain point
}

func (self *grid) Draw() {
	extraWidth := 1
	for y := self.MinY; y < self.FloorY; y++ {
		line := []byte{}
		for x := self.MinX - extraWidth; x <= self.MaxX+extraWidth; x++ {
			tile := coord{x: x, y: y}
			if tile == SandEntrypoint {
				line = append(line, '+')
				continue
			}
			currItem, ok := self.items[tile]
			if !ok {
				line = append(line, byte(ItemAir))
			} else {
				line = append(line, byte(currItem))
			}
		}
		fmt.Printf("%-3d %s\n", y, string(line))
	}
	fmt.Printf("%-3d %s\n\n", self.FloorY, strings.Repeat("#", self.MaxX+extraWidth*2-(self.MinX)+1))
}

func (self *grid) IsAir(pos coord) bool {
	if pos.y == self.FloorY {
		return false
	}

	_, ok := self.items[pos]
	if !ok {
		return true
	}
	return false
}

/// functions

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/// pre-defined values

const (
	ItemAir  item = '.'
	ItemRock item = '#'
	ItemSand item = 'o'
)

var (
	Up    = coord{x: 0, y: 1}
	Down  = coord{x: 0, y: -1}
	Right = coord{x: 1, y: 0}
	Left  = coord{x: -1, y: 0}

	SandEntrypoint = coord{500, 0}
)

func main() {
	input, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(input)

	canvas := &grid{
		items: make(map[coord]item, 0),
		// x
		MinX: 1337, // some arbitrary big number, don't really wanna import "math" just for this
		MaxX: 0,
		// y
		MinY: 0, // will always be 0
		MaxY: 0,
	}

	for scanner.Scan() {
		line := scanner.Text()
		// instructions for a single rock pattern
		rockLines := strings.Split(line, " -> ")
		//fmt.Printf("debug: %d %#v\n", len(rockLines), rockLines)

		// figure out all the coords on which the rock should be
		rockPatterns := make([]coord, 0, len(rockLines))
		for _, rl := range rockLines {
			x, _ := strconv.Atoi(rl[:strings.IndexRune(rl, ',')])   // everything before the comma
			y, _ := strconv.Atoi(rl[strings.IndexRune(rl, ',')+1:]) // everything after the comma
			if x < canvas.MinX {
				canvas.MinX = x
			} else if x > canvas.MaxX {
				canvas.MaxX = x
			}
			if y < canvas.MinY {
				canvas.MinY = y
			} else if y > canvas.MaxY {
				canvas.MaxY = y
			}
			rockPatterns = append(rockPatterns, coord{x: x, y: y})
		}

		// only actually 'draw' rocks on the grid once we parsed all pairs
		// XXX: seems inefficient...
		for i := 1; i < len(rockPatterns); i++ {
			start := rockPatterns[i-1]
			end := rockPatterns[i]

			// either of these deltas WILL be 0
			deltaX := end.x - start.x
			deltaY := end.y - start.y
			// determine in which direction to go and how many steps
			var dir coord
			var stepAmount int
			if deltaX == 0 {
				// either going up or down
				stepAmount = abs(deltaY)
				if deltaY > 0 {
					dir = Up
				} else {
					dir = Down
				}
			} else if deltaY == 0 {
				stepAmount = abs(deltaX)
				if deltaX > 0 {
					dir = Right
				} else {
					dir = Left
				}
			}
			//fmt.Printf("debug: draw rock %d times towards %v from %v to %v\n", stepAmount+1, dir, start, end)

			// actually go through all the steps, including starting one
			// but not the last one, it'd be drawn more than once - instead add it at the end, below
			for cursor, i := start, 0; i < stepAmount; i++ {
				//fmt.Println("trace: drawing", cursor)
				canvas.items[cursor] = ItemRock
				cursor.Add(dir)
			}
		}

		// because in the inner loop above, we check i < stepAmount, not i <= stepAmount
		// we need to draw a rock at the last element of last rockPattern
		canvas.items[rockPatterns[len(rockPatterns)-1]] = ItemRock
	}

	// add the floor (actually don't, just say where it should be and let GetItem handle it)
	canvas.FloorY = canvas.MaxY + 2

	// now keep adding sand until we know no more sand can fall into the cave
	// SandEntrypoint already having sand means the cave is completely full
	// XXX: is very inefficient, but is there a better way than checking every position "one by one"? maybe if I didn't use the map
	for canvas.items[SandEntrypoint] != ItemSand {
		cursor := SandEntrypoint

		// add a new sand unit
		for cameToRest := false; !cameToRest; {
			// check if we can go down

			if canvas.IsAir(coord{cursor.x, cursor.y + 1}) {
				// check if there's only void/floor below us (for getting part1 answer)
				if canvas.SandCountPart1 <= 0 && cursor.y >= canvas.MaxY {
					canvas.SandCountPart1 = canvas.SandCount
				}
				cursor.y++
				continue
			}
			// check if we can go down-left
			if canvas.IsAir(coord{cursor.x - 1, cursor.y + 1}) {
				cursor.y++
				cursor.x--
				continue
			}
			// check if we can go down-right
			if canvas.IsAir(coord{cursor.x + 1, cursor.y + 1}) {
				cursor.y++
				cursor.x++
				continue
			}

			// seems like we can't go down any further
			_, has := canvas.items[cursor]
			if !has {
				canvas.items[cursor] = ItemSand
				if cursor.x < canvas.MinX {
					canvas.MinX = cursor.x
				} else if cursor.x > canvas.MaxX {
					canvas.MaxX = cursor.x
				}
				canvas.SandCount++
				cameToRest = true
			}
		}
	}

	canvas.Draw()
	fmt.Println(canvas.SandCountPart1, "sand units before reaching floor")
	fmt.Println(canvas.SandCount, "sand units total")
}
