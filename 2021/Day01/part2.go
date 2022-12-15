package main 

import (
	"fmt";
	"os";
	"strings";
	"strconv"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	numbers := strings.Split(string(input), "\n")
	total := 0
	for i := 0; i < len(numbers) - 3; i++ {
		first, err1 := strconv.Atoi(numbers[i])
		second, err2 := strconv.Atoi(numbers[i+1])
		third, err3 := strconv.Atoi(numbers[i+2])
		fourth, err4 := strconv.Atoi(numbers[i+3])

		if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			panic(err)
		}

		if second + third + fourth > first + second + third {
			total++
		}
	}

	fmt.Println("Total of decreasings: ", total)
}