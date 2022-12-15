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

func whatIHaveToPlay(opponent string, result int) string {
	switch result {
		case 0:
			return opponent
		case 1:
			switch opponent {
				case "rock": return "paper"
				case "paper": return "scissors"
				case "scissors": return "rock"
			}
		case 2:
			switch opponent {
				case "rock": return "scissors"
				case "paper": return "rock"
				case "scissors": return "paper"
			}
		default:
			return "error"
	}
	return ""
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
		switch me {
			case "X":
				result += value(whatIHaveToPlay(translate(opponent), 2)) + 0
			case "Y":
				result += value(whatIHaveToPlay(translate(opponent), 0)) + 3
			case "Z":
				result += value(whatIHaveToPlay(translate(opponent), 1)) + 6
			default:
				panic(0)
		}
	}

	fmt.Println("The result of the strategy is: ", result);
}