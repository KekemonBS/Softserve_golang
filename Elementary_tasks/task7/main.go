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
	resString, err := Task7(flag.Bool("help", false, `triangles help`), "")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(resString.String())
	os.Exit(0)
}

func Task7(hf *bool, arg string) (strings.Builder, error) {
	var args []string
	if arg != "" {
		cmd := os.Args[0]
		args = []string{cmd, arg}
	} else {
		args = os.Args
	}

	helpFlag := hf
	flag.Parse()
	if *helpFlag {
		fmt.Println(`numspeller <number>`)
		return strings.Builder{}, nil
	}
	if len(args) != 2 {
		return strings.Builder{}, errors.New("wrong amount of arguments")
	}

	n, err := strconv.Atoi(args[1])
	if err != nil {
		return strings.Builder{}, errors.New("wrong argument passed")
	}
	return checkNums(n), nil
}

func checkNums(n int) strings.Builder {
	resString := strings.Builder{}
	for num := 0; num*num < n; num++ {
		resString.WriteString(strconv.Itoa(num))
		if (num+1)*(num+1) < n {
			resString.WriteString(", ")
		}
	}
	return resString
}
