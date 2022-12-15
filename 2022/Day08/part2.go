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

func scenicScore(trees [][]int, i int, j int) int {
	score := 1;
	addScore := 0;
	for k := j - 1; k >= 0; k-- {
		if trees[i][k] > trees[i][j] {
			addScore++;
		} else {
			score *= addScore + 1;
			break;
		}
	}

	for k := j + 1; k < len(trees[i]); k++ {
		if trees[i][k] > trees[i][j] {
			addScore++;
		} else {
			score *= addScore + 1;
			break;
		}
	}

	for k := i - 1; k >= 0; k-- {
		if trees[k][j] > trees[i][j] {
			addScore++;
		} else {
			score *= addScore + 1;
			break;
		}
	}

	for k := i + 1; k <= len(trees) - 1; k++ {
		if trees[k][j] > trees[i][j] {
			addScore++;
		} else {
			score *= addScore + 1;
			break;
		}
	}

	return score;
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

	var hihgestScenicScore int;
	for i := range trees {
		for j := range trees[i] {
			if checkBorders(trees, i, j) || checkTree(trees, i, j) {
				newScore := scenicScore(trees, i, j);
				if newScore > hihgestScenicScore {
					hihgestScenicScore = newScore;
				}
			}
		}
	}

	fmt.Println("The highest scenic score is: ", hihgestScenicScore, ".");
}
