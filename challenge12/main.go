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
	reverse := map[int]bool{}
	moves := map[int]int{}
	for i, line := range lines {
		parts := strings.Split(line, " ")
		reverse[i] = parts[0][0] == '0'
		move, _ := strconv.Atoi(parts[1])
		moves[i] = move
	}
	visited := 0
	pos := 0
	up := true
	for {
		visited++
		move, found := moves[pos]
		if !found {
			break
		}
		if reverse[pos] {
			up = !up
		}
		if up {
			pos += move
		} else {
			pos -= move
		}
	}
	fmt.Println(visited)
}
