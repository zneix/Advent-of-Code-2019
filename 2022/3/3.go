package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("failed to read input:", err)
		panic(err)
	}
	lines := strings.Fields(string(data))

	sum := 0
	for i := 0; i < len(lines); i += 3 {
		// set-like for easier checking of map's members
		duplicates := make(map[rune]struct{})

		for _, c1 := range lines[i] {
			for _, c2 := range lines[i+1] {
				if c1 == c2 {
					duplicates[c1] = struct{}{}
					break
				}
			}
		}

		for dup := range duplicates {
			if !strings.ContainsRune(lines[i+2], dup) {
				delete(duplicates, dup)
			}
		}

		// 'a'-'z' = 1-26 points
		// 'A'-'Z' = 27-52 points
		for k := range duplicates {
			switch {
			case k >= 'a' && k <= 'z':
				sum += int(k) - 96
			case k >= 'A' && k <= 'Z':
				sum += int(k) - 64 + 26
			}
			break // it should only have 1 element anyway...
		}
		//fmt.Printf("%#v %#v %d\n", i, duplicates)

	}

	fmt.Println(sum)
}
