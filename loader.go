package main

import (
	day5 "advent_of_code/day5"
	utils "advent_of_code/utils"
)

func Load() {
  input := utils.LoadData("day5/input.txt")
  day5.Solution(input);
}
