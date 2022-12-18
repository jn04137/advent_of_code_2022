package day5

import (
	"regexp"
	"strconv"
	"strings"
)

type deque []string

func (d deque) Push(e string) deque {
	return append(d, e)
}

func (d deque) Prepend(e string) deque {
	return append([]string{e}, d...)
}

func (d deque) Pop() (deque, string) {
	l := len(d)

	// Return stack with one less element;
	// Return the popped element
	return d[:l-1], d[l-1] 
}

func Move(d []deque, from int, to int, q int) (deque, deque) {
	for i := 0; i < q; i++ {
		var p string
		d[from-1], p = d[from-1].Pop()
		d[to-1] = d[to-1].Push(p)
	}

	return d[from-1], d[to-1]
}

func GetMoves(s string) []int {
	list := make([]int, 0)
	re := regexp.MustCompile("[0-9]+")
	for _, e := range re.FindAllString(s, 3) {
		n, err := strconv.Atoi(e)
		if err != nil {
			panic(err)
		}
		list = append(list, n)
	}
	return list
}

func Solution(file string) {
	// Will be used as a list of stacks
	columns := make([]deque, 9) 
	for i := range columns {
		// Each of these will be a stack
		columns[i] = make(deque, 0)
	}

	byline := strings.Split(file, "\n")
	// Line 8 is the the column name
	/* 1 5 9 13 17 21 25 29 33 */
	
	/*
		Go line-byline from 0 to 9 to collect the input data;
		This for-loop loads the initial container positions;
	*/
	for  i := 0; i < 8; i++ {
		for j, k := 1, 0; j < len(byline[i]) - 1; j, k = j+4, k+1 {
			s := string(byline[i][j])
			if s != "[" && s != "]" && !(len(strings.TrimSpace(s))==0){
				columns[k] = columns[k].Prepend(s)
			}
		}
	}

	// Executing moves in this for-loop
	for i := 11; i < len(byline); i++ {
		list := GetMoves(byline[i]):= 
		move, from, to := list[0], list[1], list[2]	
		columns[from-1], columns[to-1] = Move(columns, from, to, move)
	}
}