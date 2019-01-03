package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func intToRoman(num int) string {
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	numbers := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	str := ""
	for {
		for i := range romans {
			roman := romans[i]
			number := numbers[i]
			if num >= number {
				num -= number
				str += roman
				break
			}
		}
		if num == 0 {
			break
		}
	}
	return str
}

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	numbers := strings.Split(strings.TrimSpace(string(file)), " ")
	sum := 0
	for _, number := range numbers {
		s := 0
		n, _ := strconv.Atoi(number)
		roman := intToRoman(n)
		for _, c := range roman {
			s += int(rune(c)-rune('A')) + 1
		}
		//fmt.Printf("%d = %s = %d\n", n, roman, s)
		sum += s
	}
	fmt.Println(sum)
}
