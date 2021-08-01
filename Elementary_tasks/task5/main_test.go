package main

import (
	"flag"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func TestTask5(t *testing.T) {
	type args struct {
		hf  *bool
		arg string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Display '12345' name", args{flag.Bool("help1", false, `triangles help`), "12345"}, "дванадцять тисяч триста сорок п'ять  ", false},
		{"Display '111111111' name", args{flag.Bool("help2", false, `triangles help`), "111111111"}, "сто одинадцять мільйонів сто одинадцять тисяч сто одинадцять  ", false},
		{"Display '1' name", args{flag.Bool("help3", false, `triangles help`), "1"}, "один  ", false},
		{"Display '2' name", args{flag.Bool("help4", false, `triangles help`), "2"}, "два  ", false},
		{"Display '12' name", args{flag.Bool("help5", false, `triangles help`), "12"}, "дванадцять  ", false},
		{"Display '1000000000' name", args{flag.Bool("help6", false, `triangles help`), "1000000000"}, "один мільярд ", false},
		{"Display '1000' name", args{flag.Bool("help7", false, `triangles help`), "1000"}, "одна тисяча ", false},
		{"Display '2000' name", args{flag.Bool("help8", false, `triangles help`), "2000"}, "дві тисячі ", false},
		{"Display '235000' name", args{flag.Bool("help9", false, `triangles help`), "235000"}, "двісті тридцять п'ять тисяч ", false},
		{"Display '35000' name", args{flag.Bool("help10", false, `triangles help`), "35000"}, "тридцять п'ять тисяч ", false},
		{"Display '500' name", args{flag.Bool("help11", false, `triangles help`), "500"}, "п'ятьсот  ", false},

		{"Display help", args{flag.Bool("help12", true, `triangles help`), ""}, "", false},
		{"Malformed argument", args{flag.Bool("help", false, `triangles help`), "hello"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Task5(tt.args.hf, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Task5() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Task5() = %v, want %v", got, tt.want)
			}
		})
	}
}
