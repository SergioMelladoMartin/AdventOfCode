package main

import (
	"fmt";
	"os";
	"strings";
	"strconv";
)

func main() {
	input, err := os.ReadFile("input.txt");
	if err != nil {
		panic(err);
	}

	commands := strings.Split(string(input), "\n");
	directories := make(map[string]int);
	stackCurrentDirectories := make([]string, 0)	
	for i := 0; i < len(commands); i++ {
		words := strings.Split(commands[i], " ");
		if words[0] == "$" {
			switch words[1] {
				case "cd":
					if words[2] == ".." {
						stackCurrentDirectories = stackCurrentDirectories[:len(stackCurrentDirectories) - 1];
					} else {
						stackCurrentDirectories = append(stackCurrentDirectories, words[2]);
					}

				case "ls":
					i++;
					output := strings.Split(commands[i], " ");
					for output[0] != "$" {
						if output[0] != "dir" {
							for j := range stackCurrentDirectories {
								size, _ := strconv.Atoi(output[0]);
								directories[stackCurrentDirectories[j]] += size;
							}
						}

						if i < len(commands) - 1  {
							i++;
							output = strings.Split(commands[i], " ");
						} else {
							fmt.Println("End of file");
							break;
						}
					}
			}
		}
	}

	result := 0;
	for _, value := range directories {
		if (value < 100000) {
			result += value;
		}
	}

	fmt.Println("The result is: ", result, " bytes");
}
