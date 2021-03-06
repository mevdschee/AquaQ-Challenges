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
	sum := 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		steps, _ := strconv.Atoi(parts[0])
		size, _ := strconv.Atoi(parts[1])

		emptyRow := strings.Repeat(" ", size)
		rows := []string{}
		for y := 0; y < size+2; y++ {
			rows = append(rows, emptyRow)
		}

		for p := 2; p < len(parts); p += 2 {
			x, _ := strconv.Atoi(parts[p+1])
			y, _ := strconv.Atoi(parts[p])
			line := []byte(rows[y])
			line[x] = '#'
			rows[y+1] = string(line)
		}

		//fmt.Println(strings.Join(rows, "\n"))

		seen := map[string]int{}
		for i := 0; i < steps; i++ {
			str := strings.Join(rows, "")
			j, found := seen[str]
			if found {
				seen = map[string]int{}
				steps = i + (steps-j)%(i-j)
			}
			newRows := []string{emptyRow}
			for y := 0; y < size; y++ {
				source := strings.Join(rows[y:y+3], "")
				newRow := make([]byte, size)
				for x := 0; x < size; x++ {
					count := 0
					if x > 0 {
						if source[x+size-1] == '#' {
							count++
						}
					}
					if x < size-1 {
						if source[x+size+1] == '#' {
							count++
						}
					}
					if source[x] == '#' {
						count++
					}
					if source[x+2*size] == '#' {
						count++
					}
					if count%2 == 1 {
						newRow[x] = '#'
					} else {
						newRow[x] = '.'
					}
				}
				newRows = append(newRows, string(newRow))
			}
			newRows = append(newRows, emptyRow)
			seen[str] = i

			rows = newRows
			//fmt.Println(strings.Join(rows, "\n"))
		}
		sum += strings.Count(strings.Join(rows, ""), "#")
	}
	fmt.Println(sum)
}
