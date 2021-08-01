package main

import (
	"bufio"
	"flag"
	"strings"
	"testing"
)

func TestTask4(t *testing.T) {
	type args struct {
		hf        *bool
		argsSlice [3]string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Display help", args{flag.Bool("help1", true, `triangles help`), [3]string{"test.txt", "UNEXISTENTWORD", "NEW UNEXISTENTWORD"}}, true, false},
		{"Wrong amount of arguments", args{flag.Bool("help2", false, `triangles help`), [3]string{}}, false, true},
		{"Normal execution", args{flag.Bool("help3", false, `triangles help`), [3]string{"test.txt", "UNEXISTENTWORD", "NEW UNEXISTENTWORD"}}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Task4(tt.args.hf, tt.args.argsSlice)
			if (err != nil) != tt.wantErr {
				t.Errorf("Task4() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Task4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countWords(t *testing.T) {
	type args struct {
		scanner *bufio.Scanner
		word    string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"1 matching word", args{bufio.NewScanner(strings.NewReader("test")), "test"}, 1, false},
		{"2 matching words with obstacles", args{bufio.NewScanner(strings.NewReader("test \n\n kkkk test")), "test"}, 2, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := countWords(tt.args.scanner, tt.args.word)
			if (err != nil) != tt.wantErr {
				t.Errorf("countWords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("countWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

//File I/O requires mocking

// func Test_replaceAllInFile(t *testing.T) {
// 	type args struct {
// 		file    *os.File
// 		word    string
// 		newWord string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    string
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := replaceAllInFile(tt.args.file, tt.args.word, tt.args.newWord)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("replaceAllInFile() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if got != tt.want {
// 				t.Errorf("replaceAllInFile() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_writeToFile(t *testing.T) {
// 	type args struct {
// 		file *os.File
// 		s    string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := writeToFile(tt.args.file, tt.args.s); (err != nil) != tt.wantErr {
// 				t.Errorf("writeToFile() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
