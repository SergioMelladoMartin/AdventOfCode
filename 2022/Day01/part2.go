package main 

import (
	"fmt";
	"os";
	"strings";
	"strconv";
)

func getMin(n []int) int {
	min := 0;
	for i := range n {
		if n[i] < n[min] {
			min = i;
		}
	}
	return min;
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	numbers := strings.Split(string(input), "\n")
	var elfos []int

	j := 0;
	for i := 0; i < len(numbers) - 1; i++ {
		elfos = append(elfos, 0)
		for numbers[i] != "" {			
			kcal, err := strconv.Atoi(numbers[i])
			if err != nil {
				panic(err)
			}

			elfos[j] += kcal
			i++
		}
		j++;
	}

	maxKcal := []int{0, 0, 0};
	elfo := []int{0, 0, 0};
	for i := range elfos {
		if elfos[i] > maxKcal[getMin(maxKcal)] {
			n := getMin(maxKcal)
			maxKcal[n] = elfos[i];
			elfo[n] = i;
		}
	}

	fmt.Println("Los tres elfos con más calorías son: ", elfo[0], ", ", elfo[1], " y ", elfo[2]);
	fmt.Println("y llevan ", maxKcal[0] + maxKcal[1] + maxKcal[2], " calorías");
 }