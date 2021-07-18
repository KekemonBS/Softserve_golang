package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("luckyticket <path to file> <Moscow/Piter alg>")
		os.Exit(0)
	}

	if len(os.Args) != 3 {
		fmt.Println("Wrong amount of arguments is passed.")
		os.Exit(1)
	}

	path := os.Args[1]
	alg := os.Args[2]
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file.")
		log.Fatal(err)
		os.Exit(1)
	}

	counter := 0
	lineNumber := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if alg == "Moscow" {
			counter += moscow(scanner.Text(), lineNumber)
		} else if alg == "Piter" {
			counter += piter(scanner.Text(), lineNumber)
		}
		lineNumber++
	}
	fmt.Println("Lucky numbers : ", counter)
}

//Returns 1 if number matched to alg otherwise 0
func moscow(num string, lineNumber int) int {
	arr := toNumArr(strings.Split(num, ""), lineNumber)
	firstThreeSum := arr[0] + arr[1] + arr[2]
	lastThreeSum := arr[len(arr)-1] + arr[len(arr)-2] + arr[len(arr)-3]
	if firstThreeSum == lastThreeSum {
		return 1
	}
	return 0
}

//Returns 1 if number matched to alg otherwise 0
func piter(num string, lineNumber int) int {
	arr := toNumArr(strings.Split(num, ""), lineNumber)

	sumOfEvenNumbers := 0
	sumOfOddNumbers := 0
	for i, n := range arr {
		if i%2 == 0 {
			sumOfEvenNumbers += n
		} else {
			sumOfOddNumbers += n
		}
	}
	if sumOfOddNumbers == sumOfEvenNumbers {
		return 1
	}
	return 0
}

func toNumArr(strArr []string, lineNumber int) []int {
	var intArr []int
	for _, n := range strArr {
		num, err := strconv.Atoi(n)
		if err != nil {
			fmt.Println("Error while parsing at line : ", lineNumber+1)
			os.Exit(1)
		}
		intArr = append(intArr, num)
	}
	return intArr
}
