package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Print("chessboard [width] [height]\n")
		os.Exit(0)
	}
	width, _ := strconv.Atoi(os.Args[1])
	height, _ := strconv.Atoi(os.Args[2])
	fmt.Println(chessboard(width, height, "\033[7m \033[m"))
}

func chessboard(w int, h int, symbol string) (board string) {
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
