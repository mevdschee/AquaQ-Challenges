package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func fieldToString(field map[int]bool, size int) string {
	str := "\033[2J\033[;H"
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if field[y*1000+x] {
				str += "#"
			} else {
				str += "."
			}
		}
		str += "\n"
	}
	return str
}

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	sum := 0
	for _, line := range lines {
		field := map[int]bool{}
		parts := strings.Split(line, " ")
		steps, _ := strconv.Atoi(parts[0])
		size, _ := strconv.Atoi(parts[1])
		for p := 2; p < len(parts); p += 2 {
			x, _ := strconv.Atoi(parts[p+1])
			y, _ := strconv.Atoi(parts[p])
			field[y*1000+x] = true
		}
		for i := 0; i < steps; i++ {
			newField := map[int]bool{}
			for p := range field {
				for _, o := range []int{-1000, -1, 1, 1000} {
					pos := p + o
					y := pos / 1000
					x := pos % 1000
					if x < 0 || x >= size || y < 0 || y >= size {
						continue
					}
					if newField[pos] {
						delete(newField, pos)
					} else {
						newField[pos] = true
					}
				}
			}
			field = newField
		}
		sum += len(field)
	}
	fmt.Println(sum)
}
