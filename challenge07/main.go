package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func expectedWinRate(ratingA, ratingB float64) float64 {
	return 1 / (1 + math.Pow(10, (ratingB-ratingA)/400))
}

func adjustRatings(ratingA, ratingB float64) (float64, float64) {
	winRate := expectedWinRate(ratingA, ratingB)
	ratingWon := 20 * (1 - winRate)
	ratingA += ratingWon
	ratingB -= ratingWon
	return ratingA, ratingB
}

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	ratings := map[string]float64{}
	for _, line := range lines[1:] {
		result := strings.Split(line, ",")
		for _, player := range result[0:2] {
			ratings[player] = 1200
		}
	}
	for _, line := range lines[1:] {
		result := strings.Split(line, ",")
		player1 := result[0]
		player2 := result[1]
		scores := strings.Split(result[2], "-")
		score1, _ := strconv.Atoi(scores[0])
		score2, _ := strconv.Atoi(scores[1])
		if score1 > score2 {
			ratings[player1], ratings[player2] = adjustRatings(ratings[player1], ratings[player2])
		} else {
			ratings[player2], ratings[player1] = adjustRatings(ratings[player2], ratings[player1])
		}
	}
	min := math.MaxInt32
	max := math.MinInt32
	for _, score := range ratings {
		s := int(score)
		if s < min {
			min = s
		}
		if s > max {
			max = s
		}
	}
	fmt.Println(max - min)
}
