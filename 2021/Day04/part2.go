package main

import (
	"os"
	"strings"
	"fmt"
	"strconv"
)

type Board struct {
	numbers [5][5]int
	visited [5][5]bool
	win bool
}

func (B *Board) setNumber (x int, y int , number int) {
	B.numbers[x][y] = number
}

func (B *Board) setVisited (x int, y int) {
	B.visited[x][y] = true
}

func (B *Board) setWin () {
  B.win = true
}

func main() {
	input, _ := os.ReadFile("input.txt")	
	lines := strings.Split(string(input), "\n")

	gameNumbers := strings.Split(lines[0], ",")
	var boards []Board

	for i := 2; i < len(lines); i += 6 {
		var boardAux Board
		for j := 0; j < 5; j++ {
			splitedLine := strings.Split(lines[i + j], " ")
			for k := 0; k < 5; k++ {
				n, _ := strconv.Atoi(splitedLine[k])
				boardAux.setNumber(j, k, n)
			}
		}
		boards = append(boards, boardAux)
	}

	for i := 0; i < len(gameNumbers); i++ {
		currentNumber, _ := strconv.Atoi(gameNumbers[i]) 
		for j := 0; j < len(boards); j++ {
			for k := 0; k < 5; k++ {
				for l := 0; l < 5; l++ {
					if (boards[j].numbers[k][l] == currentNumber) { boards[j].setVisited(k, l) }
				}
			}
		}

		for j := 0; j < len(boards); j++ {
			if (!boards[j].win) {
        for k := 0; k < 5; k++ {
          nRow := 0
          nCol := 0
          for l := 0; l < 5; l++ {
            if (boards[j].visited[k][l]) { nRow++ }
            if (boards[j].visited[l][k]) { nCol++ }
          }
        
          if (nCol == 5 || nRow == 5) { boards[j].setWin() } 
        }

        nWins := 0
        for k := 0; k < len(boards); k++ {
          if (boards[k].win) { nWins++ }
        }

        if (nWins == len(boards)) {
          sum := 0
          for x := 0; x < 5; x++ {
            for y := 0; y < 5; y++ {
              if (!boards[j].visited[x][y]) { sum += boards[j].numbers[x][y] }
            }
          }
          fmt.Println(currentNumber * sum)
          os.Exit(0)
        }
      }
    }	
  }
}
