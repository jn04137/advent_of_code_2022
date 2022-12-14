package main

import (
  utils "advent_of_code/utils"
  day3 "advent_of_code/day3"
)

func Load() {
  input := utils.LoadData("day3/input.txt")
  day3.Solution(input);
}
