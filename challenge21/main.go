package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	width := 5
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	tiles := [][]int{}
	for _, line := range lines {
		numbers := []int{}
		for _, number := range strings.Split(line, " ") {
			n, _ := strconv.Atoi(number)
			numbers = append(numbers, n)
		}
		summed := []int{}
		for i := 0; i < len(numbers)-width+1; i++ {
			sum := 0
			for j := i; j < i+width; j++ {
				sum += numbers[j]
			}
			summed = append(summed, sum)
		}
		tiles = append(tiles, summed)
		//fmt.Println(summed)
	}
	fronteers := []int{}
	for pos := 0; pos < len(tiles[0]); pos++ {
		fronteers = append(fronteers, 0)
	}
	//fmt.Println("===")
	for depth := 0; depth < len(tiles); depth++ {
		newFronteers := []int{}
		for pos := 0; pos < len(fronteers); pos++ {
			max := 0
			for _, move := range []int{-1, 0, 1} {
				if pos+move < 0 || pos+move >= len(tiles[depth]) {
					continue
				}
				n := fronteers[pos+move] + tiles[depth][pos]
				if n > max {
					max = n
				}
			}
			newFronteers = append(newFronteers, max)
		}
		fronteers = newFronteers
		//fmt.Println(fronteers)
	}
	max := 0
	for pos := 0; pos < len(fronteers); pos++ {
		n := fronteers[pos]
		if n > max {
			max = n
		}
	}
	fmt.Println(max)
}
