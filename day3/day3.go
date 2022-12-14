package day3

import (
	"fmt"
	"strings"
)

// Loads the scores for each alpha-val
func LoadMap(mymap map[string]int) map[string]int {
	letters := "abcdefghijklmnopqrstuvwxyz"
	for i := range letters {
		mymap[string(letters[i])] = i + 1
	}
	for i := range letters {
		mymap[strings.ToUpper(string(letters[i]))] = len(letters) + i + 1
	}

	return mymap
}

func CheckForTwo(stringinput string, mapvals map[string]int) int {
	for i := 0; i < len(stringinput)/2; i++{
		for j := len(stringinput)/2; j < len(stringinput); j++ {
			if stringinput[i] == stringinput[j] {
				return mapvals[string(stringinput[i])]
			}
		}
	}
	return 0
}

func CheckForThree(
	s1 string, 
	s2 string, 
	s3 string, 
	mapvals map[string]int) int {
	for i := 0; i < len(s1); i++{
		for j := 0; j < len(s2); j++ {
			for k := 0; k < len(s3); k++ {
				if s1[i] == s2[j] && s1[i] == s3[k] {
					return mapvals[string(s1[i])]
				}
			}
		}
	}
	return 0
}

func Solution(input string) {
	byline := strings.Split(input, "\n")

	valuemap := make(map[string]int)
	loadedmap := LoadMap(valuemap)

	var sum int = 0
	for z := 0; z < len(byline)-1; z = z + 3 { 
		sum = sum + CheckForThree(byline[z], 
		byline[z+1], 
		byline[z+2],
		loadedmap)
	}

	fmt.Printf("Sum: %d", sum)
	// Last answer: 14061
}
