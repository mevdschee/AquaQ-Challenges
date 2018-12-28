package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func turn(dice map[string]int, faces []string){
	front := dice[faces[0]]
	dice[faces[0]]=dice[faces[1]]
	dice[faces[1]]=dice[faces[2]]
	dice[faces[2]]=dice[faces[3]]
	dice[faces[3]]=front
}

func up(dice map[string]int) {
	turn(dice, []string{"front","bottom","back","top"})
}

func down(dice map[string]int) {
	turn(dice, []string{"front","top","back","bottom"})
}

func left(dice map[string]int) {
	turn(dice, []string{"front","right","back","left"})
}

func right(dice map[string]int) {
	turn(dice, []string{"front","left","back","right"})
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	line := strings.TrimSpace(string(input))
	dice1 := map[string]int {"front": 1, "left": 2, "top": 3, "back": 6, "right": 5, "bottom": 4}
	dice2 := map[string]int {"front": 1, "left": 3, "top": 2, "back": 6, "right": 4, "bottom": 5}
	sum := 0
	for i, c := range line {
		switch c {
		case 'U':
			up(dice1)
			up(dice2)
		case 'D':
			down(dice1)
			down(dice2)
		case 'L':
			left(dice1)
			left(dice2)
		case 'R':
			right(dice1)
			right(dice2)
		}
		if dice1["front"]==dice2["front"] {
			sum += i
		}
	}
	fmt.Println(sum)
}
