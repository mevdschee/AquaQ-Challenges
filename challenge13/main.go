package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	sum := 0
	for _, line := range lines {
		maxCount := 1
		for size := 1; size < len(line); size++ {
			substrings := map[string]bool{}
			for i := 0; i <= len(line)-size; i++ {
				substrings[line[i:i+size]] = true
			}
			for substr := range substrings {
				count := 0
				s := substr
				for {
					if !strings.Contains(line, s) {
						break
					}
					s += substr
					count++
				}
				if count > maxCount {
					maxCount = count
				}
			}
			if maxCount*size > len(line) {
				break
			}
		}
		sum += maxCount
	}
	fmt.Println(sum)
}
