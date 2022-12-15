package day4

import (
	"fmt"
	"strconv"
	"strings"
)

func Solution(input string){
  finput := strings.Split(input, "\n")
  
  var sum int = 0
  for i := 0; i < len(finput)-1; i++ {
    elves := strings.FieldsFunc(string(finput[i]), Split)
    var list [4]int
    for j, val := range elves {
      num, err := strconv.Atoi(val)
      if err != nil {
        panic(err)
      }
      list[j] = num
    }

    if list[0] >= list[2] && list[0] <= list[3] {
      sum = sum + 1
    } else if list[1] >= list[2] && list[1] <= list[3] {
      sum = sum + 1 
    } else if list[2] >= list[0] && list[2] <= list[1] {
      sum = sum + 1
    } else if list[3] >= list[0] && list[3] <= list[1] {
      sum = sum + 1
    }
  }
  fmt.Printf("Current sum: %d\n", sum)
}

func Split(r rune) bool {
  return r == ',' || r == '-'
}
