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
	costs := map[string]int{}
	connections := map[string][]string{}
	for _, line := range lines {
		parts := strings.Split(line, ",")
		name1 := parts[0]
		name2 := parts[1]
		cost, _ := strconv.Atoi(parts[2])
		costs[name1+"-"+name2] = cost
		conns, found := connections[name1]
		if !found {
			conns = []string{}
		}
		connections[name1] = append(conns, name2)
	}
	sender := "TUPAC"
	receiver := "DIDDY"
	cheapest := map[string]int{}
	fronteers := map[string]int{sender: 0}
	for len(fronteers) > 0 {
		validFronteers := map[string]int{}
		for name, cost := range fronteers {
			currentCost, found := cheapest[name]
			if !found || cost < currentCost {
				cheapest[name] = cost
				validFronteers[name] = cost
			}
		}
		newFronteers := map[string]int{}
		for name1, cost := range validFronteers {
			for _, name2 := range connections[name1] {
				addedCost := costs[name1+"-"+name2]
				newFronteers[name2] = cost + addedCost
			}
		}
		fronteers = newFronteers
	}
	cost, _ := cheapest[receiver]
	fmt.Println(cost)
}
