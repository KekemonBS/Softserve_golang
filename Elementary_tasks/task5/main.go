package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var numToWord map[int]string = map[int]string{
	1:  "один",
	2:  "два",
	3:  "три",
	4:  "чотири",
	5:  "п'ять",
	6:  "шість",
	7:  "сім",
	8:  "вісім",
	9:  "де'вять",
	10: "десять",
	11: "одинадцять",
	12: "дванадцять",
	13: "тринадцять",
	14: "чотирнадцять",
	15: "п'ятнадціять",
	16: "шістнадцять",
	17: "сімнадцять",
	18: "вісімнадцять",
	19: "дев'ятнадцять",
	20: "двадцять",
	30: "тридцять",
	40: "сорок",
	50: "п'ятдесят",
	60: "шістдесят",
	70: "сімдесят",
	80: "вісімдесят",
	90: "дев'яносто",
}

var higher map[int][]string = map[int][]string{
	0: {"", ""},
	1: {"тисяч", "ж"},
	2: {"мільйон", "ч"},
	3: {"мільярд", "ч"},
}

var fractions map[int][]string = map[int][]string{
	1: {"а", ""},
	2: {"і", "и"},
	3: {"і", "и"},
	4: {"і", "и"},
}

func main() {
	fullName, err := Task5(flag.Bool("help", false, `triangles help`), "")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("|" + fullName + "|")
	os.Exit(0)
}

func Task5(hf *bool, arg string) (string, error) {
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
		return "", nil
	}
	fmt.Println(args)
	if len(args) != 2 {
		return "", errors.New("wrong number of arguments")
	}

	_, err := strconv.Atoi(args[1])
	if err != nil {
		return "", errors.New("you entered not an integer")
	}

	size := len(args[1])
	var acc []string
	var grouped []string
	for i := 0; i < size; i++ {
		d := string(args[1][i])
		if (size-i-1)%3 == 0 {
			acc = append(acc, d)
			grouped = append(grouped, strings.Join(acc, ""))
			acc = make([]string, 0)
		} else {
			acc = append(acc, d)
		}
	}

	fmt.Println(grouped)
	groupedSize := len(grouped)
	fullName := ""
	for i := 0; i < groupedSize; i++ {
		tripletName, sex := higher[groupedSize-i-1][0], higher[groupedSize-i-1][1]

		num, _ := strconv.Atoi(grouped[i])
		nameOfTriplet, tripletNamePrefix := tripletConvert(num, sex, !(i < groupedSize-1))
		if nameOfTriplet != "" {
			fullName += nameOfTriplet + tripletName + tripletNamePrefix + " "
		}
	}
	return fullName, nil
}

func tripletConvert(num int, sex string, last bool) (string, string) {
	res := ""
	fraction := ""
	hundreds := num / 100
	tens := ((num % 100) / 10) * 10
	fullTens := num % 100
	ones := num % 10

	//Form hundreds
	if hundreds != 0 {
		if numToWord[hundreds] == "один" {
			res += "сто "
		} else if numToWord[hundreds] == "два" {
			res += "двісті "
		} else if numToWord[hundreds] == "три" || numToWord[hundreds] == "чотири" {
			res += numToWord[hundreds] + "ста "
		} else {
			res += numToWord[hundreds] + "сот "
		}
	}

	if fullTens < 10 || fullTens > 19 {
		//Form tens
		if tens != 0 {
			res += numToWord[tens] + " "
		}
		//Form ones
		if ones != 0 {
			if sex == "ж" {
				if numToWord[ones] == "один" {
					res += "одна "
				} else if numToWord[ones] == "два" {
					res += "дві "
				} else {
					res += numToWord[ones] + " "
				}
			} else {
				res += numToWord[ones] + " "
			}
		}
	} else {
		res += numToWord[fullTens] + " "
	}

	//Form prefix
	if !last {
		if ones <= 4 && fractions[ones] != nil && tens == 0 {
			if sex == "ж" {
				fraction += fractions[ones][0]
			} else {
				fraction += fractions[ones][1]
			}
		} else {
			if sex == "ж" {
				fraction += ""
			} else {
				fraction += "ів"
			}
		}
	}
	return res, fraction
}
