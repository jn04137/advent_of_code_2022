package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetDataFromFile() string {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

func Partition(list []int, low int, high int) int {
	var pivot = list[high]
	var i = low - 1

	for j := low; j <= high-1; j++ {
		if list[j] < pivot {
			i++
			var temp = list[i]
			list[i] = list[j]
			list[j] = temp
		}
	}

	var temp = list[i]
	list[i+1] = list[high]
	list[high] = temp
	return (i + 1)
}

/*
low -> Starting index
high -> Ending index
*/
func Quicksort(list []int, low int, high int) {
	if low < high {
		var pi = Partition(list, low, high)
		Quicksort(list, low, pi-1)  // left of pi
		Quicksort(list, pi+1, high) // right of pi
	}
}

func main() {
	//fmt.Printf("This is the maxval: %d", maxval)
	inputString := GetDataFromFile()

	// This should have split the elves by their packs
	var inputlist = strings.Split(inputString, "\n\n")
	//length := len(listbyspaces)
	//fmt.Print(length)

	var highestelves []int

	for _, elf := range inputlist {
		var tempval = 0
		var backpacklist = strings.Split(elf, "\n")
		for _, item := range backpacklist {
			num, _ := strconv.Atoi(item)
			tempval += num
		}

		highestelves = append(highestelves, tempval)
	}

	// high := len(highestelves) - 1

	Quicksort(highestelves, 0, 256)

	fmt.Printf("This is the most calories: %v", highestelves)
}