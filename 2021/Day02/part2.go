package main

import (
	"fmt"
	"strings"
	"os"
	"strconv"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")

	x := 0
	y := 0
	aim := 0

	for _, value := range lines {
		line := strings.Split(value, " ")
		direction := line[0]
		quantity, err := strconv.Atoi(line[1])

		if err != nil {
			panic(err)
		}

		switch direction {
			case "up":
				aim -= quantity
				break

			case "down":
				aim += quantity
				break

			case "forward":
				x += quantity
				y += aim * quantity
				break
		}
	}

	fmt.Println("Solution: ", x * y)
}
