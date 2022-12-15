package main 

import (
	"fmt";
	"os";
	"strings";
	"strconv";
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	numbers := strings.Split(string(input), "\n")
	total := 0
	for i := 0; i < len(numbers) - 1; i++ {
		first, err1 := strconv.Atoi(numbers[i])
		second, err2 := strconv.Atoi(numbers[i+1])

		if err1 != nil || err2 != nil {
			panic(err)
		}
		
		if second > first {
			total++
		}
	}
 
	fmt.Println("Total of decreasings: ", total)
}