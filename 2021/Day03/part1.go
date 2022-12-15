package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	
	numbers := strings.Split(string(input), "\n")
	numberSize := len(numbers[0]) - 1
	
	var numberOfOnes []int 
	for i := 0; i < numberSize; i++ {
		numberOfOnes = append(numberOfOnes, 0)
	}

	for i := 0; i < len(numbers); i++ {
		splitedLine := strings.Split(numbers[i], "")
		for j := 0; j < numberSize; j++ {
			if splitedLine[j] == "1" {
				numberOfOnes[j]++
			}
		}
	}

	gamma, epsilon := "", ""
	for i := 0; i < len(numberOfOnes); i++ {
		if numberOfOnes[i] > len(numbers) / 2 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	gammaInt, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonInt, _ := strconv.ParseInt(epsilon, 2, 64)
	fmt.Println(gammaInt * epsilonInt)
}