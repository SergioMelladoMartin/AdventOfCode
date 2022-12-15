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

	maxKcal := 0;
	elfo := 0;
	for i := range elfos {
		if elfos[i] > maxKcal {
			maxKcal = elfos[i];
			elfo = i;
		}
	}

	fmt.Println("Elfo ", elfo, " lleva: ", maxKcal, " calorías");
 }