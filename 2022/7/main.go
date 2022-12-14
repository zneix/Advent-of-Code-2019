package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	SMALL_DIR_SIZE = 100_000
	DISK_SIZE      = 70_000_000
	REQUIRED_SPACE = 30_000_000
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file) // fancy way of reading inputs while skipping newlines. kinda like C's fgets

	dirSizes := make(map[string]int)
	cwd := "/" // because of this, we skip the first line of input below - imo, it's not worth handling smth that will never happen
	scanner.Scan()
	for scanner.Scan() {
		args := strings.Fields(scanner.Text())
		//fmt.Println("debug:", args)
		// determine what kind of input does current line represent
		if args[0] == "$" {
			// a command entered by the user
			if args[1] == "cd" {
				cwd = filepath.Join(cwd, args[2])
				//fmt.Printf("changed to %q %s (size %d)\n", args[2], cwd, dirSizes[cwd])
			}
		} else {
			// terminal's output, assume it's from ls (since no other command gives us any output)
			if args[0] != "dir" {
				fileSize, err := strconv.Atoi(args[0])
				if err != nil {
					panic(err) // won't occur unless input is malformed
				}
				dirSizes[cwd] += fileSize
				// XXX: add file's size recursively; can this be optimized?
				dir := strings.Clone(cwd) // not sure Clone is necessary, but better safe than sorry
				numParents := len(strings.FieldsFunc(cwd, func(r rune) bool { return r == '/' }))
				//fmt.Printf("cwd+: %s %d, %d parents\n", cwd, dirSizes[cwd], numParents)
				//for i := numParents - 1; i >= 0; i-- {
				for i := 0; i < numParents; i++ {
					dir = filepath.Join(dir, "..")
					//dir = dir[:strings.LastIndexByte(dir, '/')] // TODO: fix this and use it instead of filepath.Join
					//fmt.Printf("recur+: %d/%d %q %d\n", numParents-i, numParents, dir, fileSize)
					dirSizes[dir] += fileSize
				}

				//fmt.Printf("ls: %#v %d\n", args, fileSize)
			} else {
				//fmt.Printf("ls: directory %s (size %d)\n", args[1], dirSizes[filepath.Join(cwd, args[1])])
				//fmt.Printf("ls: directory %s (size %d)\n", args[1], dirSizes[cwd+"/"+args[1]])
			}
		}
	}

	smallDirSum := 0
	dirToDelete := "/" // just some value that for sure won't be the final result
	neededSpace := REQUIRED_SPACE - (DISK_SIZE - dirSizes["/"])
	for k, v := range dirSizes {
		if v <= SMALL_DIR_SIZE {
			smallDirSum += v
		}

		if v >= neededSpace && v < dirSizes[dirToDelete] {
			//fmt.Printf("Found new smallest dir: %d < %d (%s < %s)\n", v, dirSizes[dirToDelete], k, dirToDelete)
			dirToDelete = k
		}
	}

	// part 1
	fmt.Println(smallDirSum)

	// part 2
	//fmt.Printf("Disk used: %d\nDisk free: %d\nNeed: %d\n", dirSizes["/"], DISK_SPACE-dirSizes["/"], neededSpace)
	fmt.Println(dirSizes[dirToDelete])
}
