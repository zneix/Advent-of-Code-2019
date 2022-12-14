package main

import (
	"fmt"
	"os"
)

const n = 14 // amount of unique elements we're after

// TODO: investigate sliding window algorithm for this..
func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	for i := n - 1; i < len(data)-1; i++ {
		subset := data[i-n+1 : i+1]
		seen := make(map[byte]struct{})
		died := false

		for _, v := range subset {
			if _, ok := seen[v]; !ok {
				seen[v] = struct{}{}
			} else {
				died = true
				break
			}
		}
		if !died {
			fmt.Println(i + 1)
			break
		}
	}
}
