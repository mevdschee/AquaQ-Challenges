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
	boughtMilk := []int{}
	freshCereal := 0
	freshMilk := 0
	for i, line := range lines[1:] {
		for j:=-5;j<0;j++ {
			if i+j>=0 && boughtMilk[i+j]>=100 {
				if freshCereal>=100 {
					freshCereal -= 100
					boughtMilk[i+j] -= 100
					break
				}
			}
		}
		groceries := strings.Split(line, ",")
		milk, _ := strconv.Atoi(groceries[1])
		cereal, _ := strconv.Atoi(groceries[2])
		boughtMilk = append(boughtMilk, milk)
		freshCereal += cereal
		freshMilk = 0
		for j:=-4;j<=0;j++ {
			if i+j>=0 {
				freshMilk+=boughtMilk[i+j]
			}
		}
		//fmt.Printf("%d: %d %d\n",i+1,freshMilk,freshCereal)
	}
	fmt.Println(freshMilk + freshCereal)
}
