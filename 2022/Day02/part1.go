package main

import (
	"fmt";
	"os";
	"strings";
)

func translate (symbol string) string {
	if symbol == "A" || symbol == "X" {
		return "rock"
	} else if symbol == "B" || symbol == "Y" {
		return "paper"
	} else if symbol == "C" || symbol == "Z" {
		return "scissors"
	} else {
		return "error"
	}
}

func rockPaperScissors(player1 string, player2 string) int {
	if player1 == player2 {
		return 0
	} 

	switch player1 {
		case "rock":
			if player2 == "paper" {
				return 2
			} else {
				return 1
			}

		case "paper":
			if player2 == "scissors" {
				return 2
			} else {
				return 1
			}

		case "scissors":
			if player2 == "rock" {
				return 2
			} else {
				return 1
			}
		
		default:
			return -1
	}
}

func value(input string) int {
	if input == "rock" {
		return 1
	} else if input == "paper" {
		return 2
	} else if input == "scissors" {
		return 3
	} else {
		return 0
	}
}


func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	plays := strings.Split(string(input), "\n")
	result := 0;
	for i := range plays {
		splitedPlay := strings.Split(plays[i], " ")
	    opponent := splitedPlay[0];
		me := splitedPlay[1];
		winner := rockPaperScissors(translate(me), translate(opponent));
		switch winner {
			case 0: result += 3 + value(translate(me));
			case 1: result += 6 + value(translate(me));
			case 2: result += 0 + value(translate(me));
			default: panic(0);
		}
	}

	fmt.Println("The result of the strategy is: ", result);
}