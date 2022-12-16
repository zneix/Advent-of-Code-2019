package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type coord struct {
	x, y int
}

type numRange struct {
	start, end int
}

// I could've reused 'coord', but that'd feel too confusing
type setRange struct {
	start, end   int
	start2, end2 int
	hasBreak     bool
}

type grid map[int][]numRange

func tuningFreq(x, y int) int {
	return x*4000000 + y
}

// cba typecasting with equivalents from math package
func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// XXX: BUGGED, doesn't remove any elements from the array that have been shadowed later on
func sumRange(ogRanges []numRange) []numRange {
	//var result []numRange
	result := make([]numRange, 0)

	for _, r := range ogRanges {
		r.start = max(r.start, 0)
		r.end = min(LIMIT, r.end)
		isInFinal := false

		// check against ranges constructed during this function's run
		for i, resR := range result {
			//fmt.Printf("debug: c%v res[%d]%v\n", r, i, resR)
			if r.start <= resR.end && r.end >= resR.start {
				// r starts before resR ends or
				// r ends after resR starts
				//fmt.Printf("trace1: %v %v; %v %v\n", r.start, resR.end, r.end, resR.start)
				result[i].start = min(r.start, resR.start)
				result[i].end = max(r.end, resR.end)
				isInFinal = true
				break
			} else if r.start == resR.end+1 {
				// r is right after resR
				//fmt.Printf("trace2: %v %v; %v %v\n", r.start, resR.end+1, r.end, result[i].end)
				result[i].end = r.end
				isInFinal = true
				break
			} else if r.end == resR.end-1 {
				// r is right before resR
				//fmt.Printf("trace3: %v %v; %v %v\n", r.end, resR.end-1, r.start, result[i].start)
				result[i].start = r.start
				isInFinal = true
				break
			}
		}
		//fmt.Printf("trace4 %-5t %v fres%v\n", isInFinal, r, result)

		// current range isn't in the resulting array of ranges yet, let's add it
		if !isInFinal {
			result = append(result, r)
		}
	}

	return result
}

const LIMIT = 4000000

// bro...
func (self grid) pushSafe(y int, v numRange) {
	if y < 0 || y > LIMIT {
		return
	}

	if (self)[y] == nil {
		(self)[y] = make([]numRange, 0)
	}

	(self)[y] = append((self)[y], v)
}

func main() {
	log.SetFlags(log.Flags() | log.Lmicroseconds)

	input, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}
	lines := strings.FieldsFunc(string(input), func(r rune) bool { return r == '\n' })

	// boundaries or smth
	canvas := make(grid)
	for li, line := range lines {
		s, b := coord{}, coord{}
		fmt.Sscanf(line, `Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d`, &s.x, &s.y, &b.x, &b.y)

		// now process this...
		deltaX := abs(b.x - s.x)
		deltaY := abs(b.y - s.y)
		dist := deltaX + deltaY

		// good lord, what is happening in there...
		// XXX: remove this aurora borealis (or maybe it's actually good, idk)
		log.Printf("dank bruteforce: %d\n", li)
		for i := 1; i <= dist; i++ {
			hashStart := s.x - (dist - i)
			hashEnd := s.x + (dist - i)

			canvas.pushSafe(s.y-i, numRange{hashStart, hashEnd})
			canvas.pushSafe(s.y+i, numRange{hashStart, hashEnd})
		}
		canvas.pushSafe(s.y, numRange{s.x - dist, s.x + dist})
	}

	// try to sum ranges now
	for y, row := range canvas {
		//fmt.Printf("row %d(%d): %v\n", y, len(row), row)
		sum := sumRange(sumRange(sumRange(row))) // XXX: very, Very, VERY dirty temp solution to sumRange's bug
		//fmt.Printf("sum %d(%d): %v\n\n", y, len(sum), sum)
		if len(sum) > 1 {
			log.Printf("yay %d(%d): %v\n", y, len(sum), sum)
			fmt.Println(tuningFreq(sum[0].end+1, y))
			return // we can exit already
		}
	}

	log.Println("xd", len(canvas))
}
