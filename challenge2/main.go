package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"strconv"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	line := strings.TrimSpace(string(input))
	numberStrings := strings.Split(line," ")
	numbers := []int{}
	for _,str := range numberStrings {
		number, _ := strconv.Atoi(str)
		numbers = append(numbers, number)
	}
	for {
		memory := map[int]int{}
		var next []int
		for i,number := range numbers {
			pos, found := memory[number]
			if found {
				next = append(numbers[:pos],numbers[i:]...)
				break
			}
			memory[number] = i
		}
		if next==nil {
			break
		}
		numbers = next
	}
	sum := 0
	for _,number := range numbers {
		sum+=number
	}
	fmt.Println(sum)
}