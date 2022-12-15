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

func coincidence3 (input1 string, input2 string, input3 string) string {
	for i := range input1 {
		for j := range input2 {
			for k := range input3 {
				if input1[i] == input2[j] && input2[j] == input3[k] {
					return string(input1[i])
				}
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
	for i := 0; i < len(bags); i += 3 { 
		result += value(coincidence3(bags[i], bags[i + 1], bags[i + 2]));
	}

	fmt.Println("The result is: ", result);
}