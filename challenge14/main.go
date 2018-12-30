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

	cardStrings := "6 17 34 50 68\n10 21 45 53 66\n5 25 36 52 69\n14 30 33 54 63\n15 23 41 51 62"
	coordinates := map[int]int{}
	size := 0
	for y, cardString := range strings.Split(cardStrings, "\n") {
		size++
		for x, num := range strings.Split(cardString, " ") {
			c := y*10 + x
			n, _ := strconv.Atoi(num)
			coordinates[n] = c
		}
	}

	sum := 0
	for _, line := range lines {
		draws := 0
		checked := map[int]bool{}
		for _, num := range strings.Split(line, " ") {
			draws++
			n, _ := strconv.Atoi(num)
			c, found := coordinates[n]
			if !found {
				continue
			}
			checked[c] = true

			done := false
			// horizontal
			for y := 0; y < size; y++ {
				count := 0
				for x := 0; x < size; x++ {
					if checked[y*10+x] {
						count++
					}
				}
				if count == size {
					done = true
					break
				}
			}
			// vertical
			for x := 0; x < size; x++ {
				count := 0
				for y := 0; y < size; y++ {
					if checked[y*10+x] {
						count++
					}
				}
				if count == size {
					done = true
					break
				}
			}
			// diagonal
			for d := 0; d < 2; d++ {
				count := 0
				for y := 0; y < size; y++ {
					x := y
					if d == 1 {
						x = size - 1 - y
					}
					if checked[y*10+x] {
						count++
					}
				}
				if count == size {
					done = true
					break
				}
			}

			if done {
				break
			}
		}
		sum += draws
	}
	fmt.Println(sum)
}
