package main

import (
	"fmt";
	"os";
	"strings";
	"strconv";
)

func areUnderhanded(firstRange []int, secondRange []int) bool {
	if firstRange[1] < secondRange[0] || secondRange[1] < firstRange[0] {
		return false;
	}
	return true;
}

func main() {
	input, err := os.ReadFile("input.txt");
	if err != nil {
		panic(err);
	}

	ranges := strings.Split(string(input), "\n");
	count := 0;
	for _, value := range ranges {
		firstRange := make([]int, 2);
		secondRange := make([]int, 2);

		splittedRanges := strings.Split(value, ",");
		for j, value := range strings.Split(splittedRanges[0], "-") {
			firstRange[j], _ = strconv.Atoi(value);
		}

		for j, value := range strings.Split(splittedRanges[1], "-") {
			secondRange[j], _ = strconv.Atoi(value);
		}
		
		if (areUnderhanded(firstRange, secondRange)) {
			count++;
		}
	}

	fmt.Println("The inputs are ", count, " underhanded");
}