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

type letter struct {
	name   string
	width  float64
	height float64
}

func main() {
	err := Task2(flag.Bool("help", false, `canfit help`), [2]string{"", ""})
	if err != nil {
		fmt.Printf("err : %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func Task2(hf *bool, proposedDims [2]string) error {
	helpFlag := hf
	flag.Parse()

	if *helpFlag {
		fmt.Println(`canifit - program will prompt you to describe dimmensions of two letters
				 and compace if you can fit one of them in other one`)
		return nil
	}
	L1dims, L1Name := fillLetter("L1", proposedDims[0])
	L2dims, L2Name := fillLetter("L2", proposedDims[1])

	L1, err := convertLetter(L1dims, L1Name)
	if err != nil {
		return err
	}

	L2, err := convertLetter(L2dims, L2Name)
	if err != nil {
		return err
	}

	L1inL2 := compareTwoLetters(L1, L2)
	L2inL1 := compareTwoLetters(L2, L1)
	fmt.Println("\nCan you fit L1 in L2 ?", L1inL2)
	fmt.Println("Can you fit L2 in L1 ?", L2inL1)
	return nil
}

func fillLetter(name string, proposed string) (string, string) {
	if len(proposed) != 0 {
		return proposed, name
	}

	fmt.Printf("\nPlease fill in %s dimensions: \n", name)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Width : ")
	scanner.Scan()
	w := scanner.Text()

	fmt.Print("Height : ")
	scanner.Scan()
	h := scanner.Text()

	return w + "," + h, name
}

func convertLetter(dimensions string, name string) (letter, error) {
	whslice := strings.Split(strings.Replace(dimensions, " ", "", -1), ",")
	l1width, err := strconv.ParseFloat(whslice[0], 32)
	if err != nil {
		return letter{}, errors.New("error parsing width")
	}

	l1height, err := strconv.ParseFloat(whslice[1], 32)
	if err != nil {
		return letter{}, errors.New("error parsing height")

	}

	return letter{name, l1width, l1height}, nil
}

func compareTwoLetters(L1 letter, L2 letter) bool {
	if L1.width*L1.height > L2.width*L2.height {
		return false
	}
	if !((L1.height < L2.height && L1.width < L2.width) || (L1.width < L2.height && L1.height < L2.width)) {
		return false
	}
	return true
}
