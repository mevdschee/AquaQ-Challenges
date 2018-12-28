package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	line := strings.TrimSpace(string(input))
	roomStrings := strings.Split("  ##  \n #### \n######\n######\n #### \n  ##  ","\n")
	ymax := len(roomStrings)-1
	xmax := len(roomStrings[0])-1
	rooms := map[int]bool{}
	for y,str := range roomStrings {
		for x,c := range str {
			if c=='#' {
				rooms[y*10+x] = true
			}
		}
	}
	x := 2
	y := 0
	sum := 0
	for _,c := range line {
		nx := x
		ny := y
		switch c {
		case 'U':
			if ny > 0 {
				ny -= 1
			}
		case 'D':
			if ny < ymax {
				ny += 1
			}
		case 'L':
			if nx > 0 {
				nx -= 1
			}
		case 'R':
			if nx < xmax {
				nx += 1
			}
		}
		_, found := rooms[ny * 10 + nx]
		if !found {
			nx = x
			ny = y
		}
		x = nx
		y = ny
		sum += y + x
	}
	fmt.Println(sum)
}