package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	resString, err := Task7(flag.Bool("help", false, `triangles help`), [2]string{})
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(resString)
	os.Exit(0)
}

func Task7(hf *bool, arg [2]string) (string, error) {
	var args []string
	if arg[0] != "" && arg[1] != "" {
		cmd := os.Args[0]
		args = []string{cmd, arg[0], arg[1]}
	} else {
		args = os.Args
	}

	helpFlag := hf
	flag.Parse()
	if *helpFlag {
		fmt.Println(`fib <from> <to>`)
		return "", nil
	}
	fmt.Println(args)
	if len(args) != 3 {
		return "", errors.New("wrong amount of arguments")
	}

	from, err := strconv.Atoi(args[1])
	if err != nil {
		return "", errors.New("entered wrong first argument")
	}
	to, err := strconv.Atoi(args[2])
	if err != nil {
		return "", errors.New("entered wrong second argument")
	}

	fibList := fromToFibList(from, to)
	return strings.Join(fibList, ", "), nil
}

func fromToFibList(from, to int) []string {
	fibFrom := fib()
	next := 0
	var fibList []string
	for next < to {
		next = fibFrom()
		if next < to && next > from-1 {
			fibList = append(fibList, strconv.Itoa(next))
		}
	}
	return fibList
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
