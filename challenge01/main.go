package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	line := strings.TrimSpace(string(input))
	re := regexp.MustCompile(`[^0-9a-f]`)
	line = re.ReplaceAllString(line, `0`)
	for len(line)%3 != 0 {
		line += string('0')
	}
	answer := ""
	for i := 0; i < len(line); i += len(line) / 3 {
		answer += line[i : i+2]
	}
	fmt.Println(answer)
}
