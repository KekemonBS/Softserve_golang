package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	count, err := Task6(flag.Bool("help", false, `triangles help`), [2]string{})
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Lucky numbers : ", count)
	os.Exit(0)
}

func Task6(hf *bool, argsSlice [2]string) (int, error) {
	var args []string
	if argsSlice[0] != "" && argsSlice[1] != "" {
		cmd := os.Args[0]
		args = []string{cmd, argsSlice[0], argsSlice[1]}
	} else {
		args = os.Args
	}

	helpFlag := hf
	flag.Parse()
	if *helpFlag {
		fmt.Println(`luckyticket <path to file> <Moscow/Piter alg>`)
		return 0, nil
	}

	if len(args) != 3 {
		return 0, errors.New("wrong amount of arguments is passed")
	}

	path := args[1]
	alg := args[2]
	counter, errorCounter := countLucky(path, alg)
	if errorCounter != nil {
		return 0, errorCounter
	}
	return counter, nil
}

func countLucky(path string, alg string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, errors.New("error opening file")
	}

	counter := 0
	lineNumber := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) < 6 {
			return 0, errors.New("ticket number must be at least 6 characters wide, error at line : " + fmt.Sprint(lineNumber+1))
		}

		if alg == "Moscow" {
			match, err := moscow(scanner.Text(), lineNumber)
			if err != nil {
				return 0, err
			}
			counter += match
		} else if alg == "Piter" {
			match, err := piter(scanner.Text(), lineNumber)
			if err != nil {
				return 0, err
			}
			counter += match
		}
		lineNumber++
	}
	return counter, nil
}

//Returns 1 if number matched to alg otherwise 0
func moscow(num string, lineNumber int) (int, error) {
	arr, err := toNumArr(strings.Split(num, ""), lineNumber)
	if err != nil {
		return 0, err
	}

	firstThreeSum := arr[0] + arr[1] + arr[2]
	lastThreeSum := arr[len(arr)-1] + arr[len(arr)-2] + arr[len(arr)-3]
	if firstThreeSum == lastThreeSum {
		return 1, nil
	}
	return 0, nil
}

//Returns 1 if number matched to alg otherwise 0
func piter(num string, lineNumber int) (int, error) {
	arr, err := toNumArr(strings.Split(num, ""), lineNumber)
	if err != nil {
		return 0, err
	}

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
		return 1, nil
	}
	return 0, nil
}

func toNumArr(strArr []string, lineNumber int) ([]int, error) {
	var intArr []int
	for _, n := range strArr {
		num, err := strconv.Atoi(n)
		if err != nil {
			return []int{}, errors.New("Error while parsing at line : " + fmt.Sprint(lineNumber+1))
		}
		intArr = append(intArr, num)
	}
	return intArr, nil
}
