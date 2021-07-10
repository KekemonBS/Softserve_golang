package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/KekemonBS/Softserve_golang/day1/tasks"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Choose task [1-4] : ")
	scanner.Scan()
	taskNum := scanner.Text()
	switch taskNum {
	case "1":
		{
			fmt.Println("--------Task1--------")

			fmt.Print("Chessboard width : ")
			scanner.Scan()
			width, err := strconv.Atoi(scanner.Text())
			if err != nil || width <= 0 {
				fmt.Println("Excepted integer greater than zero")
				os.Exit(1)
			}

			fmt.Print("Chessboard height : ")
			scanner.Scan()
			height, err := strconv.Atoi(scanner.Text())
			if err != nil || height <= 0 {
				fmt.Println("Excepted integer greater than zero")
				os.Exit(1)
			}

			fmt.Print("Chessboard symbol : ")
			scanner.Scan()
			symbol := scanner.Text()

			fmt.Println()
			fmt.Println(tasks.Task1(width, height, symbol))
		}
	case "2":
		{
			fmt.Println("--------Task2--------")

			fmt.Print("Type in array of integers separated by comma. \nArray :")
			scanner.Scan()
			arr := scanner.Text()

			_, q := tasks.Task2(arr)
			fmt.Println("Quantity of positive eneven numbers :", q)
		}
	case "3":
		{
			fmt.Println("--------Task3--------")
			fmt.Print("Enter credit card number : ")
			scanner.Scan()
			creditCardNumber := scanner.Text()
			fmt.Println()
			isValid, lastFourDigits := tasks.Task3(creditCardNumber)
			if isValid {
				fmt.Println("Credit card number is valid.\nLast four digits :", lastFourDigits)
			} else {
				fmt.Println("Credit card number is invalid.")
			}
		}
	case "4":
		{

			fmt.Println("--------Task4--------")
			//init

			fmt.Print("Enter how many Fibonacci numbers do you need : ")
			scanner.Scan()
			n, err := strconv.Atoi(scanner.Text())
			if err != nil || n <= 0 {
				fmt.Println("Cant generate less than zero numbers.")
				os.Exit(1)
			}
			fmt.Println()
			fib := tasks.Task4()
			var fibList []string
			for i := 0; i < n; i++ {
				fibList = append(fibList, strconv.Itoa(fib()))
			}
			fmt.Println(strings.Join(fibList, ", "))

		}
	}
}
