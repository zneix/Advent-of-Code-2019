package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("failed to read input:", err)
		panic(err)
	}
	lines := strings.Fields(string(data))

	sum := 0
	for _, line := range lines {
		// tbh I could've just used len() and avoid having to import unicode/utf8 package (since input only has a-zA-Z)
		// but it's better to use the method below to check string length
		// so I'm just using that to keep up a good habit
		hl := utf8.RuneCountInString(line) / 2
		first, second := line[hl:], line[:hl]

		var duplicate byte
		// check for duplicate being uninitialized to save iterations
		// doing this manually with 2 loops to kinda make it bit more efficient(?)
		// and eliminate overhead that comes from the underlying strings.Index(s1, s2) call
		for i := 0; i < hl && duplicate == 0; i++ {
			for j := 0; j < hl; j++ {
				if first[i] == second[j] {
					duplicate = first[i]
					break
				}
			}
		}

		// 'a'-'z' = 1-26 points
		// 'A'-'Z' = 27-52 points
		switch {
		case duplicate >= 'a' && duplicate <= 'z':
			sum += int(duplicate) - 96
		case duplicate >= 'A' && duplicate <= 'Z':
			sum += int(duplicate) - 64 + 26
		}
		//fmt.Printf("%#v %s %d\n", i, string(duplicate))

	}

	fmt.Println(sum)
}
