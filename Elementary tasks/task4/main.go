package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {
	var word string
	var newWord string
	if len(os.Args) == 1 {
		fmt.Println(`The program must receive arguments for input when starting:
		<file path> <string for counting>
		<file path> <search string> <replacement string>`)
		os.Exit(0)
	}

	path := os.Args[1]
	if len(os.Args) >= 3 {
		word = os.Args[2]
	}
	if len(os.Args) == 4 {
		newWord = os.Args[3]
	}
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	counter := 0
	matchRegex := "\\b" + word + "\\b"
	for scanner.Scan() {
		match, _ := regexp.MatchString(matchRegex, scanner.Text())
		if match {
			counter++
		}
	}

	if len(os.Args) == 4 {
		data, _ := ioutil.ReadFile(path)
		//fmt.Println("All text:", data)
		re := regexp.MustCompile(word)
		s := re.ReplaceAllString(string(data), newWord)
		//fmt.Println("All text after:", s)
		_ = file.Truncate(0)
		file.WriteAt([]byte(s), 0)
		file.Close()
		fmt.Printf("\nAll occurances of word %s replaced by %s\n", word, newWord)
	} else {
		fmt.Printf("\nWord %s found %d times\n", word, counter)
	}
}
