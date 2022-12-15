package main

import (
	"os"
	"strings"
	"fmt"
	"strconv"
)

func getMax(a int , b int) int {
	if a > b { return a }
	return b
}

func getMin(a int, b int) int {
	if a < b { return a }
	return b
}

func plus1(a int, b int) int {
	if a < b { return 1 }
	return -1
}

func main() {
	input, err := os.ReadFile("input.txt")	
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	var board [1000][1000]int

	for i := range lines {
		data      := strings.Split(string(lines[i]), ",")
		first, _  := strconv.Atoi(data[0])
		second, _ := strconv.Atoi(data[1])
		third, _  := strconv.Atoi(data[2])
		fourth, _ := strconv.Atoi(data[3])

		if first == third {
			for j := getMin(second, fourth); j <= getMax(second,	fourth); j++ {
				board[j][first]++
			} 
		} else if second == fourth {
			for j := getMin(first, third); j <= getMax(first, third); j++ {
				board[second][j]++
			}
	  } else {
			x := first
			y := second

			for {
				board[x][y]++
				x += plus1(x, third)
				y += plus1(y, fourth)
				if x == third || y == fourth { break }
			}
		}
	}

	count := 0
	for i := range board {
		for j := range board[i] {
			if board[i][j] > 1 { count++ }
		}
	}
	fmt.Println("Result:", count)
}