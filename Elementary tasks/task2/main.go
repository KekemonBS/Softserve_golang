package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

type letter struct {
	name   string
	width  float64
	height float64
}

func main() {
	helpFlag := flag.Bool("help", false, `canfit help`)
	flag.Parse()

	if *helpFlag {
		fmt.Println(`canifit - program will prompt you to describe dimmensions of two letters
				 and compace if you can fit one of them in other one`)
		os.Exit(1)
	}
	L1 := fillLetter("L1")
	L2 := fillLetter("L2")
	L1inL2 := compareTwoLetters(L1, L2)
	L2inL1 := compareTwoLetters(L2, L1)
	fmt.Println("\nCan you fit L1 in L2 ?", L1inL2)
	fmt.Println("Can you fit L2 in L1 ?", L2inL1)
}

func fillLetter(name string) letter {
	fmt.Printf("\nPlease fill in %s dimensions: \n", name)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Width : ")
	scanner.Scan()
	l1width, _ := strconv.ParseFloat(scanner.Text(), 32)

	fmt.Print("Height : ")
	scanner.Scan()
	l1height, _ := strconv.ParseFloat(scanner.Text(), 32)

	return letter{name, l1width, l1height}
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
