package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("sqlessthan <number>")
		os.Exit(0)
	}
	if len(os.Args) != 2 {
		fmt.Println("Wrong amount of arguments.")
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Wrong argument passed.")
		os.Exit(1)
	}

	resString := strings.Builder{}
	for num := 0; num*num < n; num++ {
		resString.WriteString(strconv.Itoa(num))
		if (num+1)*(num+1) < n {
			resString.WriteString(", ")
		}
	}
	fmt.Println(resString.String())
}
