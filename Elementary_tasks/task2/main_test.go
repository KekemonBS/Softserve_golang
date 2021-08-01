package main

import (
	"errors"
	"flag"
	"reflect"
	"testing"
)

func Test_compareTwoLetters(t *testing.T) {
	type args struct {
		L1 letter
		L2 letter
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"L1 fits in L2 ", args{letter{"L1", 1.0, 1.0}, letter{"L2", 3.0, 3.0}}, true},
		{"L1 not fits in L2 ", args{letter{"L1", 3.0, 3.0}, letter{"L2", 1.0, 1.0}}, false},
		{"L1 same as L2", args{letter{"L1", 3.0, 3.0}, letter{"L2", 3.0, 3.0}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareTwoLetters(tt.args.L1, tt.args.L2); got != tt.want {
				t.Errorf("compareTwoLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertLetter(t *testing.T) {
	type args struct {
		dimensions string
		name       string
	}
	type want struct {
		letter letter
		err    error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		//{"Normal values", args{"5, 10", "L1"}, want{letter{"L1", 5, 10}, nil}},
		{"Wrong width", args{"4k, 8", "L1"}, want{letter{}, errors.New("error parsing width")}},
		{"Wrong height", args{"3, 8j", "L1"}, want{letter{}, errors.New("error parsing height")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertLetter(tt.args.dimensions, tt.args.name)
			if !reflect.DeepEqual(got, tt.want.letter) && err.Error() != tt.want.err.Error() {
				t.Errorf("compareTwoLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTask2(t *testing.T) {
	type args struct {
		hf           *bool
		proposedDims [2]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"No help flag, no failure", args{flag.Bool("help1", false, `canfit help`), [2]string{"1 ,2", "1, 1"}}, false},
		{"Help flag, no filure", args{flag.Bool("help2", true, `canfit help`), [2]string{"", ""}}, false},
		{"No help flag, failure (L1 err)", args{flag.Bool("help3", false, `canfit help`), [2]string{"1 k 21kh ", "2, 2"}}, true},
		{"No help flag, failure (L2 err)", args{flag.Bool("help4", false, `canfit help`), [2]string{"2, 2", "1 k 21kh "}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Task2(tt.args.hf, tt.args.proposedDims); (err != nil) != tt.wantErr {
				t.Errorf("Task2() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
