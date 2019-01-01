package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func isPalindrome(s string) bool {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	return s == string(runes[n:])
}

func findPalindromeTimes() map[string]int {
	palindromeTimes := map[string]int{}
	for h := 0; h < 24; h++ {
		for m := 0; m < 60; m++ {
			for s := 0; s < 60; s++ {
				str := fmt.Sprintf("%02d:%02d:%02d", h, m, s)
				if isPalindrome(str) {
					palindromeTimes[str] = h*3600 + m*60 + s
				}
			}
		}
	}
	return palindromeTimes
}

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	palindromeTimes := findPalindromeTimes()
	sum := 0
	for _, line := range lines {
		var h, m, s int
		fmt.Sscanf(line, "%02d:%02d:%02d", &h, &m, &s)
		time := h*3600 + m*60 + s
		min := math.MaxInt32
		for _, t := range palindromeTimes {
			for day := -1; day <= 1; day++ {
				dt := t - time + day*24*3600
				if dt < 0 {
					dt *= -1
				}
				if dt < min {
					min = dt
				}
			}
		}
		sum += min
	}
	fmt.Println(sum)
}
