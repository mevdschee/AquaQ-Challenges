package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readAlphabet(filename string) map[rune][]string {
	file, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(file), "\n")
	letters := map[rune][]string{}
	for i, line := range lines {
		ch := rune(int('A') + i/6)
		if i%6 == 0 {
			letters[ch] = []string{}
		}
		letters[ch] = append(letters[ch], line)
	}
	return letters
}

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	line := strings.TrimSpace(string(file))

	letters := readAlphabet("asciialphabet.txt")

	result := []string{"", "", "", "", "", ""}
	for pos, ch := range line {
		max := 0
		for i, line := range letters[ch] {
			left := strings.TrimRight(result[i], " ")
			right := strings.TrimLeft(line, " ")
			length := len(left) + len(right)
			if length > max {
				max = length
			}
		}
		if pos > 0 {
			max++
		}
		for i, line := range letters[ch] {
			left := strings.TrimRight(result[i], " ")
			right := strings.TrimLeft(line, " ")
			spaces := strings.Repeat(" ", max-len(left)-len(right))
			result[i] = left + spaces + right
		}
	}

	sum := 0
	for _, line := range result {
		sum += strings.Count(line, " ")
	}
	fmt.Println(sum)
}
