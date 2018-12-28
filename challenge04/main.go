package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	prime, _ := strconv.Atoi(strings.TrimSpace(string(input)))
	sum := 0
	for i := 1; i < prime; i++ {
		if gcd(prime, i) == 1 {
			sum += i
		}
	}
	fmt.Println(sum)
}
