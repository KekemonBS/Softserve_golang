package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	err := Task1(flag.Bool("help", false, `canfit help`), []string{})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func Task1(hf *bool, args []string) error {
	var widthStr string
	var heightStr string
	if len(args) != 0 {
		widthStr = args[0]
		heightStr = args[1]
	} else {
		widthStr = os.Args[1]
		heightStr = os.Args[2]
	}

	helpFlag := hf
	flag.Parse()
	if *helpFlag {
		fmt.Println("chessboard <width> <height>")
		return nil
	}

	width, err := strconv.Atoi(widthStr)
	if err != nil {
		return errors.New("error parsing width")
	}
	height, err := strconv.Atoi(heightStr)
	if err != nil {
		return errors.New("error parsing height")
	}
	fmt.Println(Chessboard(width, height, "\033[7m \033[m"))
	return nil
}

func Chessboard(w int, h int, symbol string) (board string) {
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if i%2 == 0 {
				board += symbol + " "
			} else {
				board += " " + symbol
			}
		}
		board += "\n"
	}
	return board
}
