package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type coord struct {
	x, y int
}

func (self *coord) tuningFreq() int {
	return self.x*4000000 + self.y
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

const WANTED_Y = 2000000

func main() {
	input, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.FieldsFunc(string(input), func(r rune) bool { return r == '\n' })
	minX := math.MaxInt
	maxX := math.MinInt
	beaconsAtWanted := make(map[coord]struct{}) // set of unique beacons at WANTED_Y

	for _, line := range lines {
		s, b := coord{}, coord{}
		fmt.Sscanf(line, `Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d`, &s.x, &s.y, &b.x, &b.y)

		// now process this...
		deltaX := abs(b.x - s.x)
		deltaY := abs(b.y - s.y)
		dist := deltaX + deltaY

		deltaWanted := abs(s.y - WANTED_Y)
		howLongX := dist - deltaWanted

		if b.y == WANTED_Y {
			beaconsAtWanted[b] = struct{}{}
		}

		if howLongX > 0 {
			//fmt.Printf("line: %v %v\n%d\n%d\n", s, b, deltaY, deltaWanted)
			//fmt.Printf("debug: %d <= x <= %d\n", s.x-howLongX, s.x+howLongX)
			if s.x-howLongX < minX {
				minX = s.x - howLongX
			}
			if s.x+howLongX > maxX {
				maxX = s.x + howLongX
			}
			//hashSets = append(hashSets, set{s.x - howLongX, s.x + howLongX})
		}
	}
	// XXX: just assuming there won't be any "empty" spaces in between all the ranges...
	fmt.Println(abs(minX) + abs(maxX) + 1 - len(beaconsAtWanted))
}
