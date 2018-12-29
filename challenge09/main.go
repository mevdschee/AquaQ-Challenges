package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"strconv"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	total := big.NewInt(1)
	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		total = total.Mul(total, big.NewInt(int64(num)))
	}
	fmt.Println(total)
}
