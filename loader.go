package main

import (
  utils "advent_of_code/utils"
  //day3 "advent_of_code/day3"
  day4 "advent_of_code/day4"
)

func Load() {
  input := utils.LoadData("day4/input.txt")
  day4.Solution(input);
}
