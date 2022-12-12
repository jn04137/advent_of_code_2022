package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

/*
=== Inputs ===
First column is the opponent and the second column is
what your response should be.

A - Rock
B - Paper
C - Scissors

X - Rock
Y - Paper
Z - Scissors

Points
Rock 			- 1 point
Paper 		- 2 points
Scissors 	- 3 points
*/

/*
	Personal Note, you can read the list list line by line and
	compare two letters within each line, or you can put every
	letter into a stock and compare the letters as they come
	out of the stack.
*/

func LoadData(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

func ListByGame(input string) []string {
	return strings.Split(input, "\n")
}

func ListByInput(input string) []string {
	return strings.Split(input, "\\s")
}

func InputComparison(opponent string, myinput string) int {
	inputmap := make(map[string]int)
	inputmap["A"] = 1 // Rock
	inputmap["B"] = 2 // Paper
	inputmap["C"] = 3 // Scissors

	/*
		C Y = 	
	*/

	if myinput == "X" {
		// need to lose
		if inputmap[opponent] == 1 {
			return 3	
		} else if inputmap[opponent] == 2 {
			return 1
		} else if inputmap[opponent] == 3 {
			return 2
		}
	} else if myinput == "Y" {
		// need to draw
		if inputmap[opponent] == 1 {
			return 1 + 3
		} else if inputmap[opponent] == 2 {
			return 2 + 3
		} else if inputmap[opponent] == 3 {
			return 3 + 3
		}
	} else {
		// need to win
		if inputmap[opponent] == 1 {
			return 2 + 6
		} else if inputmap[opponent] == 2 {
			return 3 + 6
		} else if inputmap[opponent] == 3 {
			return 1 + 6
		}
	}

	return inputmap[myinput];
	// list of 200 object
	// 1 step
	// log base(10) 0, 1, 2, 3...
	// log base (2) 0/1
}

func main() {
	var score int = 0;
	var input string = LoadData("input.txt")

	games := ListByGame(input)

	fmt.Printf("Number of games: %d\n", len(games))

	//fmt.Printf("First element: %v\n", games[0])
	//split1ele := strings.Split(games[0], "\\s")
	//fmt.Printf("Split first element: %v\n", split1ele)
	for i := 0; i < len(games) - 1; i++ {
		element := string(games[i])
		fmt.Printf("Current element: %s\n", string(element[0]))
		fmt.Printf("Current element: %s\n", string(element[2]))
		score = score + InputComparison(string(element[0]), string(element[2]))
	}

	fmt.Printf("Total score: %d", score)
}
