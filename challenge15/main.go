package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readWords(filename string) map[string]bool {
	words := map[string]bool{}
	file, _ := ioutil.ReadFile("words.txt")
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	for _, line := range lines {
		words[strings.TrimSpace(line)] = true
	}
	return words
}

func findTransformations(word1 string, words map[string]bool) []string {
	transformations := []string{}
	letters := "abcdefghijklmnopqrstuvwxyz"
	for i := range word1 {
		for j := range letters {
			word2 := word1[:i] + letters[j:j+1] + word1[i+1:]
			if words[word2] && word2 != word1 {
				transformations = append(transformations, word2)
			}
		}
	}
	return transformations
}

func findShortestPath(sourceWord, targetWord string, words map[string]bool) (string, bool) {
	paths := map[string]string{sourceWord: sourceWord}
	seen := map[string]bool{}
	for {
		newPaths := map[string]string{}
		for path, currentWord := range paths {
			if !seen[currentWord] {
				seen[currentWord] = true
				for _, nextWord := range findTransformations(currentWord, words) {
					newPath := path + " " + nextWord
					if nextWord == targetWord {
						return newPath, true
					}
					newPaths[newPath] = nextWord
				}
			}
		}
		paths = newPaths
		if len(paths) == 0 {
			break
		}
	}
	return "", false
}

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	words := readWords("words.txt")
	product := 1
	for _, line := range lines {
		transformation := strings.Split(line, ",")
		source := transformation[0]
		target := transformation[1]
		path, found := findShortestPath(source, target, words)
		//fmt.Println(path)
		if found {
			product *= len(strings.Split(path, " "))
		}
	}
	fmt.Println(product)
}
