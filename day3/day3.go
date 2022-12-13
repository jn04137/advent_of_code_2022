package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func LoadData(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

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
				if s1[i] == s2[j] && s1[i] == s[k] {
					return mapvals[string(s1[i])]
				}
			}
		}
	}
	return 0
}

func main() {
	input := LoadData("input.txt")
	byline := strings.Split(input, "\n")

	valuemap := make(map[string]int)
	loadedmap := LoadMap(valuemap)

	//var match string = ""
	var sum int = 0
	// For every input line
	for z := 0; z < len(byline); z++ { 
		// For every element in first half of string
			sum = sum + CheckForTwo(byline[z], loadedmap)
	}

	fmt.Printf("Sum: %d", sum)
	// Last answer: 14061
}