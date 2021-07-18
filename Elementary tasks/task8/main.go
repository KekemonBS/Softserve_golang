package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("fib <from> <to>")
		os.Exit(0)
	}
	
	if len(os.Args) != 3 {
		fmt.Println("Wrong amount of arguments.")
		os.Exit(1)
	}
	
	from, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Entered wrong first argument.")
		os.Exit(1)
	}
	to, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Entered wrong second argument.")
		os.Exit(1)
	}

	fibFrom := fib()
	next := 0
	var fibList []string
	for next < to {
		next = fibFrom()
		if next < to && next > from-1 {
				fibList = append(fibList, strconv.Itoa(next))
		}
	}
	fmt.Println(strings.Join(fibList, ", "))
}

func fib() func() (n int) {
	var prev int
	var curr int
	var next int

	return func() (fibNextNum int) {
		next = curr + prev

		if prev == 0 {
			curr = 1
		}

		prev = curr
		curr = next
		return next
	}
}
