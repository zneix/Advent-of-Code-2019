package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var linePattern = regexp.MustCompile(`^Valve ([A-Z]+) has flow rate=(\d+); tunnels? leads? to valves? (.+)$`)

func main() {
	input, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}
	lines := strings.FieldsFunc(string(input), func(r rune) bool { return r == '\n' })

	for _, line := range lines {
		matches := linePattern.FindStringSubmatch(line)

		valve := matches[1]
		flowRate, _ := strconv.Atoi(matches[2])
		nextValves := strings.Split(matches[3], ", ")

		fmt.Printf("%d %q %v\n", flowRate, valve, nextValves)
	}
}
