package main

import (
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
	0: []string{"", ""},
	1: []string{"тисяч", "ж"},
	2: []string{"мільйон", "ч"},
	3: []string{"мільярд", "ч"},
}

var fractions map[int][]string = map[int][]string{
	1: []string{"а", ""},
	2: []string{"і", "и"},
	3: []string{"і", "и"},
	4: []string{"і", "и"},
}

func main() {

	if len(os.Args) == 1 {
		fmt.Println("numspeller <number>")
		os.Exit(0)
	} else if len(os.Args) != 2 {
		fmt.Println("Wrong number of arguments.")
		os.Exit(1)
	}

	_, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("You entered not an integer.")
		os.Exit(1)
	}

	size := len(os.Args[1])
	var acc []string
	var grouped []string
	for i := 0; i < size; i++ {
		d := string(os.Args[1][i])
		if (size-i-1)%3 == 0 {
			acc = append(acc, d)
			grouped = append(grouped, strings.Join(acc, ""))
			acc = make([]string, 0)
		} else {
			acc = append(acc, d)
		}
	}

	fmt.Println(grouped)
	//resString := strings.Builder{}

	groupedSize := len(grouped)
	fullName := ""
	for i := 0; i < groupedSize; i++ {
		tripletName, sex := higher[groupedSize-i-1][0], higher[groupedSize-i-1][1]

		num, _ := strconv.Atoi(grouped[i])
		nameOfTriplet, tripletNamePrefix := tripletConvert(num, sex, !(i < groupedSize-1))
		fullName += nameOfTriplet + tripletName + tripletNamePrefix + " "
	}
	fmt.Println(fullName)
}

func tripletConvert(num int, sex string, last bool) (string, string) {
	res := ""
	fraction := ""
	hundreds := num / 100
	tens := ((num % 100) / 10) * 10
	fullTens := num % 100
	ones := num % 10
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
	if numToWord[fullTens] == "" {
		if tens != 0 {
			res += numToWord[tens] + " "
		}

		if ones != 0 {
			if sex == "ж" {
				if numToWord[ones] == "один" {
					res += numToWord[ones] + "а "
				} else if numToWord[ones] == "два" {
					res += "дві" + " "
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
	if !last {
		if fractions[ones] != nil && numToWord[fullTens] == "" {
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
