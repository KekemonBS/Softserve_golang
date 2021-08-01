package main

import (
	"flag"
	"testing"
)

type wantGet struct {
	inW int
	inH int
	out string
}

func Test_Chessboard3x1(t *testing.T) {
	wantGet := wantGet{
		inW: 3,
		inH: 1,
		out: "\033[7m \033[m \033[7m \033[m \033[7m \033[m \n",
	}
	t1result := Chessboard(wantGet.inW, wantGet.inH, "\033[7m \033[m")

	if t1result != wantGet.out {
		t.Errorf("got %q, wanted %q", t1result, wantGet.out)
	}
}

func Test_Chessboard1x3(t *testing.T) {
	wantGet := wantGet{
		inW: 1,
		inH: 3,
		out: "\033[7m \033[m \n \033[7m \033[m\n\033[7m \033[m \n",
	}
	t1result := Chessboard(wantGet.inW, wantGet.inH, "\033[7m \033[m")

	if t1result != wantGet.out {
		t.Errorf("got %q, wanted %q", t1result, wantGet.out)
	}
}

func Test_Chessboard2x2(t *testing.T) {
	wantGet := wantGet{
		inW: 2,
		inH: 2,
		out: "\033[7m \033[m \033[7m \033[m \n \033[7m \033[m \033[7m \033[m\n",
	}
	t1result := Chessboard(wantGet.inW, wantGet.inH, "\033[7m \033[m")

	if t1result != wantGet.out {
		t.Errorf("got %q, wanted %q", t1result, wantGet.out)
	}
}

func TestTask1(t *testing.T) {
	type args struct {
		hf   *bool
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Help display", args{flag.Bool("help1", true, `canfit help`), []string{}}, false},
		{"Basic chessboard", args{flag.Bool("help2", false, `canfit help`), []string{"2", "2"}}, false},
		{"Wrong height", args{flag.Bool("help3", false, `canfit help`), []string{"2", "2ha"}}, true},
		{"Wrong width", args{flag.Bool("help4", false, `canfit help`), []string{"2ha", "2"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Task1(tt.args.hf, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("Task1() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
