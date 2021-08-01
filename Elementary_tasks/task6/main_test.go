package main

import (
	"flag"
	"reflect"
	"testing"
)

func TestTask6(t *testing.T) {
	type args struct {
		hf        *bool
		argsSlice [2]string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"Display help", args{flag.Bool("help2", true, `triangles help`), [2]string{}}, 0, false},
		{"Wrong amouunt of arguments eror", args{flag.Bool("help3", false, `triangles help`), [2]string{}}, 0, true},
		{"Count from file with Piter alg twice", args{flag.Bool("help4", false, `triangles help`), [2]string{"test.txt", "Piter"}}, 2, false},
		{"Error in file", args{flag.Bool("help5", false, `triangles help`), [2]string{"test_withErr.txt", "Piter"}}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Task6(tt.args.hf, tt.args.argsSlice)
			if (err != nil) != tt.wantErr {
				t.Errorf("Task6() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Task6() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countLucky(t *testing.T) {
	type args struct {
		path string
		alg  string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"Matched by Moscow alg three times", args{"test.txt", "Moscow"}, 3, false},
		{"Matched by Piter alg twice", args{"test.txt", "Piter"}, 2, false},
		{"Error opening file", args{"testkkk.txt", "Piter"}, 0, true},
		{"Wrong ticket ( < 5 chars )", args{"test_withErr.txt", "Piter"}, 0, true},
		{"Wrong number of ticket", args{"test_withErrInNum.txt", "Piter"}, 0, true},
		{"Wrong number of ticket", args{"test_withErrInNum.txt", "Moscow"}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := countLucky(tt.args.path, tt.args.alg)
			if (err != nil) != tt.wantErr {
				t.Errorf("countLucky() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("countLucky() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_moscow(t *testing.T) {
	type args struct {
		num        string
		lineNumber int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"Matched", args{"132231", 1}, 1, false},
		{"Did not matched", args{"101111", 1}, 0, false},
		{"Wrong arg", args{"k", 1}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := moscow(tt.args.num, tt.args.lineNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("moscow() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("moscow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_piter(t *testing.T) {
	type args struct {
		num        string
		lineNumber int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"Matched", args{"123321", 1}, 1, false},
		{"Did not matched", args{"123621", 1}, 0, false},
		{"Wrong arg", args{"k", 1}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := piter(tt.args.num, tt.args.lineNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("piter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("piter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toNumArr(t *testing.T) {
	type args struct {
		strArr     []string
		lineNumber int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{"Test string to arr", args{[]string{"7", "2", "4"}, 1}, []int{7, 2, 4}, false},
		{"Test err", args{[]string{"7k", "2", "4"}, 1}, []int{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := toNumArr(tt.args.strArr, tt.args.lineNumber); !reflect.DeepEqual(got, tt.want) || (err != nil) != tt.wantErr {
				t.Errorf("toNumArr() = %v, want %v", got, tt.want)
			}
		})
	}
}
