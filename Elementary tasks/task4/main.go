package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var word string
	var newWord string

	path := os.Args[1]
	if len(os.Args) == 3 {
		word := os.Args[2]
	} 
	if len(os.Args) == 4 {
		newWord := os.Args[3]	
	} 
	fileHandle, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	fmt.Printf("\nPlease select mode [Substitute - s/]: \n")
	scannerAns := bufio.NewScanner(os.Stdin)
	
	counter := 0
	for scanned := scanner.Scan() {
		if regexp.MatchString("\\bword\\b", scanned) {
			counter++
		}
	}

	if len(os.Args) == 4 {
			s := re.ReplaceAllString(word, newWord)
			os.
			fmt.Printf("All occurances of word %s replaced by %s", word, newWord)		
	}
	
	fmt.Printf("Word %s found %d times", word, counter)
	
}
