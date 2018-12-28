package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func adders(result string, parts, left int) int {
	if parts == 1 {
		result += strconv.Itoa(left)
		sum := 0
		for _, c := range result {
			if c == '1' {
				sum += 1
			}
		}
		fmt.Println(result)
		return sum
	}
	sum := 0
	for i := 0; i <= left; i++ {
		sum += adders(result+strconv.Itoa(i)+" ", parts-1, left-i)
	}
	return sum
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	line := strings.Split(strings.TrimSpace(string(input)), " numbers which sum to ")
	parts, _ := strconv.Atoi(line[0])
	target, _ := strconv.Atoi(line[1])
	fmt.Println(adders("", parts, target))
}
