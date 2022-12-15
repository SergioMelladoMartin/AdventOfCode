package main

import (
	"fmt";
	"os";
	"strings";
)

func value(input string) int {
	if input >= "a" && input <= "z" {
		return int(input[0]) - 96
	}
	return int(input[0]) - 38
}

func coincidence (input1 string, input2 string) string {
	for i := range input1 {
		for j := range input2 {
			if input1[i] == input2[j] {
				return string(input1[i])
			}
		}
	}
	return "error"
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	bags := strings.Split(string(input), "\n");
	result := 0;
	for i := range bags {
		firstCompartment := bags[i][0: len(bags[i]) / 2];
		secondCompartment := bags[i][len(bags[i]) / 2: len(bags[i])];
		result += value(coincidence(firstCompartment, secondCompartment));
	}

	fmt.Println("The result is: ", result);
}