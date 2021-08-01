package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
)

func main() {
	_, err := Task4(flag.Bool("help", false, `triangles help`), [3]string{})
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func Task4(hf *bool, argsSlice [3]string) (bool, error) {
	var word string
	var newWord string

	var args []string
	if argsSlice[0] != "" && argsSlice[1] != "" && argsSlice[2] != "" {
		cmd := os.Args[0]
		args = []string{cmd, argsSlice[0], argsSlice[1], argsSlice[2]}
	} else {
		args = os.Args
	}

	helpFlag := hf
	flag.Parse()
	if *helpFlag {
		fmt.Println(`The program must receive arguments for input when starting:
		<file path> <string for counting>
		<file path> <search string> <replacement string>`)
		return true, nil
	}

	if len(args) > 4 {
		return false, errors.New("wrong amount of arguments")
	}

	path := args[1]
	if len(args) >= 3 {
		word = args[2]
	}
	if len(args) == 4 {
		newWord = args[3]
	}
	file, errOpen := os.OpenFile(path, os.O_RDWR, 0666)
	if errOpen != nil {
		return false, errOpen
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	counter, errCount := countWords(scanner, word)
	if errCount != nil {
		return false, errCount
	}

	if len(args) == 4 {
		s, errReplace := replaceAllInFile(file, word, newWord)
		if errReplace != nil {
			return false, errReplace
		}
		errWrite := writeToFile(file, s)
		if errWrite != nil {
			return false, errWrite
		}
		fmt.Printf("\nAll occurances of word %s replaced by %s\n", word, newWord)
	} else {
		fmt.Printf("\nWord %s found %d times\n", word, counter)
	}
	return true, nil
}

func countWords(scanner *bufio.Scanner, word string) (int, error) {
	counter := 0
	matchRegex := "\\b" + word + "\\b"
	for scanner.Scan() {
		match, err := regexp.MatchString(matchRegex, scanner.Text())
		if err != nil {
			return 0, errors.New("error while matching regex")
		}
		if match {
			counter++
		}
	}
	return counter, nil
}

func replaceAllInFile(file *os.File, word string, newWord string) (string, error) {
	fi, err1 := file.Stat()
	if err1 != nil {
		return "", err1
	}
	data := make([]byte, fi.Size())
	file.Seek(0, 0)
	_, err2 := file.Read(data)
	if err2 != nil {
		return "", errors.New("error reading file")
	}
	re := regexp.MustCompile(word)
	s := re.ReplaceAllString(string(data), newWord)
	return s, nil
}

func writeToFile(file *os.File, s string) error {
	err3 := file.Truncate(0)
	if err3 != nil {
		return errors.New("error truncating file")
	}
	file.WriteAt([]byte(s), 0)
	return nil
}
