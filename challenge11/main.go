package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	tiles := map[int]int{}
	for _, line := range lines[1:] {
		coordinates := strings.Split(line, ",")
		lx, _ := strconv.Atoi(coordinates[0])
		ly, _ := strconv.Atoi(coordinates[1])
		ux, _ := strconv.Atoi(coordinates[2])
		uy, _ := strconv.Atoi(coordinates[3])
		for y := ly; y < uy; y++ {
			for x := lx; x < ux; x++ {
				count, _ := tiles[y*100+x]
				tiles[y*100+x] = count + 1
			}
		}
	}
	neighbours := []int{-100, 100, -1, 1}
	tileCount := 0
	for {
		changed := false
		tileCount = 0
		for coordinates, count := range tiles {
			if count != 2 {
				continue
			}
			tileCount++
			for _, offset := range neighbours {
				count, found := tiles[coordinates+offset]
				if found && count != 2 {
					tiles[coordinates+offset] = 2
					changed = true
				}
			}
		}
		if !changed {
			break
		}
	}
	fmt.Println(tileCount)
}
