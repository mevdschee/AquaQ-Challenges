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
		field := map[int]bool{}
		parts := strings.Split(line, " ")
		steps, _ := strconv.Atoi(parts[0])
		size, _ := strconv.Atoi(parts[1])
		for p := 2; p < len(parts); p += 2 {
			x, _ := strconv.Atoi(parts[p+1])
			y, _ := strconv.Atoi(parts[p])
			field[y*1000+x] = true
		}

		emptyRow := strings.Repeat(" ", size)
		rows := []string{emptyRow}
		for y := 0; y < size; y++ {
			row := ""
			for x := 0; x < size; x++ {
				if field[y*1000+x] {
					row += "#"
				} else {
					row += "."
				}
			}
			rows = append(rows, row)
		}
		rows = append(rows, emptyRow)

		//fmt.Println(strings.Join(rows, "\n"))

		fieldSeen := map[string]int{}
		rowSeen := map[string][]byte{}
		for i := 0; i < steps; i++ {
			str := strings.Join(rows, "")
			j, found := fieldSeen[str]
			if found {
				fieldSeen = map[string]int{}
				steps = i + (steps-j)%(i-j)
			}
			newRows := []string{emptyRow}
			for y := 0; y < size; y++ {
				source := strings.Join(rows[y:y+3], "")
				newRow, found := rowSeen[source]
				if !found {
					newRow = make([]byte, size)
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
					rowSeen[source] = newRow
				}
				newRows = append(newRows, string(newRow))
			}
			newRows = append(newRows, emptyRow)
			fieldSeen[str] = i

			rows = newRows
			//fmt.Println(strings.Join(rows, "\n"))
		}
		sum += strings.Count(strings.Join(rows, "\n"), "#")
	}
	fmt.Println(sum)
}
