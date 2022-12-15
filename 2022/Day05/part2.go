package main

import (
	"os";
	"strings";
	"fmt";
	"strconv";
)

func insertLine(matrix [][]string, line string) [][]string{
	line = strings.ReplaceAll(line, "[", "");
	line = strings.ReplaceAll(line, "]", "");
	line = strings.ReplaceAll(line, "  ", " ");
	var newRow []string;
	for i := 0; i < len(line); i += 2 {
		newRow = append(newRow, string(line[i]));
	}
	return append(matrix, newRow);
}

func concatElements(matrix [][]string) string {
	result := ""
	for j := 0; j < len(matrix[0]); j++ {
		for i := 0; i < len(matrix); i++ {
		    if matrix[i][j] != " " {
				result += matrix[i][j];
				break;
			}
		}
	}

	return result;
}

func moveContainers(matrix [][]string, move int, from int, to int) [][]string {
	var elements []string;
	for i := 0; i < len(matrix) && len(elements) < move; i++ {
		if matrix[i][from] != " " {
			elements = append(elements, matrix[i][from])
			matrix[i][from] = " "
		}
	}

	for i := len(matrix) - 1; i >= 0 && len(elements) > 0; i-- {
		if matrix[i][to] == " " {
			matrix[i][to] = elements[len(elements) - 1]
			elements = elements[:len(elements) - 1]
		}
	}

	if len(elements) >= 0 {
		var newMatrix [][]string;
		var newRow []string;
		for i := 0; i < len(elements); i++ {
			for j := 0; j < len(matrix[0]); j++ {
				if j == to {
					newRow = append(newRow, elements[i])
				} else {
					newRow = append(newRow, " ")
				}
			}
			newMatrix = append(newMatrix, newRow)
			newRow = nil
		}
		newMatrix = append(newMatrix, matrix...)
		return newMatrix;
	} else {
		return matrix;
	}
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n");
	var containers [][]string
	lineWhenStartMovements := 0;
	for i, line := range lines {
		if string(line[1]) != "1" {
			containers = insertLine(containers, line);
		} else {
			lineWhenStartMovements = i
			break;
		}
	}

	for i := lineWhenStartMovements + 2; i < len(lines); i++ {
		line := strings.Split(lines[i], " ");
		move, _ := strconv.Atoi(line[1]);
		from, _ := strconv.Atoi(line[3]);
		to, _ := strconv.Atoi(line[5]);
		containers = moveContainers(containers, move, from - 1, to - 1);
	}

	fmt.Println("The result is: ", concatElements(containers));
}