package main

import (
	"fmt"
	"os"
	"strings"
)

func areAllDifferent(characters []string) bool {
	for i := range characters {
		for j := range characters {
			if i != j && strings.Contains(characters[i], characters[j]) {
				return false
			}
		}
	}
	return true
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	characters := make([]string, 14)
	result := 0
	for i, letter := range string(input) {
		characters[i%14] = string(letter)
		if areAllDifferent(characters) {
			result = i + 1;
			break
		}
	}

	fmt.Println("The result is:", result)
}