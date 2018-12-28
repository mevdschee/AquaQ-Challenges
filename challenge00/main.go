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
	answer := ""
	letters := []string{" ", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		number, _ := strconv.Atoi(parts[0])
		count, _ := strconv.Atoi(parts[1])
		len := len(letters[number])
		answer += string(letters[number][(count-1)%len])
	}
	fmt.Println(answer)
}
