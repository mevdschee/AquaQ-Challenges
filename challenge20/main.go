package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	cards := strings.Split(strings.TrimSpace(string(file)), " ")
	won := 0
	sum := 0
	aces := 0
	for _, card := range cards {
		num, err := strconv.Atoi(card)
		if err != nil {
			if card == "A" {
				aces++
				num = 11
			} else {
				num = 10
			}
		}
		//fmt.Print(card + " ")
		sum += num
		if sum > 21 && aces > 0 {
			aces--
			sum -= 10
		}
		if sum == 21 {
			won++
		}
		if sum >= 21 {
			//fmt.Print("= ")
			//fmt.Println(sum)
			sum = 0
			aces = 0
		}
	}
	//fmt.Println()
	fmt.Println(won)
}
