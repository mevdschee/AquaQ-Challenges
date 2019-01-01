package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	starts := map[string]string{}
	max := 0
	maxCountry := ""
	maxFrom := ""
	maxUntil := ""
	for _, line := range lines[1:] {
		fields := strings.Split(line, ",")
		date := fields[0]
		teams := fields[1:3]
		scores := fields[3:5]
		for i := range teams {
			if scores[i] == "0" {
				if _, found := starts[teams[i]]; !found {
					starts[teams[i]] = date
				}
			} else {
				if _, found := starts[teams[i]]; found {
					layout := "2006-01-02"
					t1, _ := time.Parse(layout, starts[teams[i]])
					t2, _ := time.Parse(layout, date)
					duration := int(t2.Sub(t1).Hours() / 24)
					if duration > max {
						max = duration
						maxCountry = teams[i]
						maxFrom = starts[teams[i]]
						maxUntil = date
					}
					delete(starts, teams[i])
				}
			}
		}
	}
	maxFrom = strings.Replace(maxFrom, "-", "", 2)
	maxUntil = strings.Replace(maxUntil, "-", "", 2)
	fmt.Println(maxCountry, maxFrom, maxUntil)
}
