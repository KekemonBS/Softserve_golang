package tasks

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Task2(arr string) (numArray []int, q int) {
	arr = strings.ReplaceAll(arr, " ", "")
	intStrings := strings.Split(arr, ",")
	for _, num := range intStrings {
		integer, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println(num, " is not an integer.")
			os.Exit(1)
		}
		if integer > 0 && integer%2 == 0 {
			q++
		}
		numArray = append(numArray, integer)
	}
	return numArray, q
}
