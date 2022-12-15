package main 

import (
	"fmt";
	"os";
	"strings";
	"strconv";
)

func checkBorders(trees [][]int, i int, j int) bool {
	return i == 0 || i == len(trees) - 1 || j == 0 || j == len(trees[i]) - 1;
}

func checkTree(trees [][]int, i int, j int) bool {
	visible := [4]bool{true, true, true, true};
	for k := j - 1; k >= 0; k-- {
		if trees[i][k] >= trees[i][j] {
			visible[0] = false;
		}
	}

	for k := j + 1; k < len(trees[i]); k++ {
		if trees[i][k] >= trees[i][j] {
			visible[1] = false;
		}
	}

	for k := i - 1; k >= 0; k-- {
		if trees[k][j] >= trees[i][j] {
			visible[2] = false;
		}
	}

	for k := i + 1; k <= len(trees) - 1; k++ {
		if trees[k][j] >= trees[i][j] {
			visible[3] = false;
		}
	}

	for i := range visible {
		if visible[i] {
			return true;
		}
	}

	return false
}

func main() {
	input, err := os.ReadFile("input.txt");
	if err != nil {
		panic(err);
	}

	trees := make([][]int, 0);
	rows := strings.Split(string(input), "\n");
	for i := range rows {
		var newRow []int;
		for j := range rows[i] {
			n, _ := strconv.Atoi(string(rows[i][j]));
			newRow = append(newRow, n);
		}
		trees = append(trees, newRow);
	}

	var numberOfVisibleTrees int;

	for i := range trees {
		for j := range trees[i] {
			if checkBorders(trees, i, j) {
				numberOfVisibleTrees++;
			} else {
				if checkTree(trees, i, j) {
					numberOfVisibleTrees++;
				}
			}
		}
	}

	fmt.Println("The number of visible trees is", numberOfVisibleTrees, ".");
}
