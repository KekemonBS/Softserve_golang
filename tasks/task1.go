package tasks

func Task1(w int, h int, symbol string) (board string) {
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
