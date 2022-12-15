package main

import (
	"fmt"
	"strings"
	"os"
	"strconv"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	x := 0
	y := 0

	for _, value := range lines {
		line := strings.Split(value, " ")
		direction := line[0]
		quantity, err := strconv.Atoi(line[1])
		if err != nil {
			panic(err)
		}

		switch direction {
			case "up":
				y -= quantity
				break

			case "down":
				y += quantity
				break

			case "forward":
				x += quantity
				break
		}
	}	

	fmt.Println("Solution: ", x * y)
}