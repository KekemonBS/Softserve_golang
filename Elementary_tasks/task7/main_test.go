package main

import (
	"flag"
	"reflect"
	"strings"
	"testing"
)

func TestTask7(t *testing.T) {
	type args struct {
		hf  *bool
		arg string
	}
	tests := []struct {
		name    string
		args    args
		want    strings.Builder
		wantErr bool
	}{
		{"Display help", args{flag.Bool("help1", true, `triangles help`), ""}, strings.Builder{}, false},
		{"Squares lesser than 10", args{flag.Bool("help2", false, `triangles help`), "10"}, getBuilderWithText("0, 1, 2, 3"), false},
		{"Wrong argument", args{flag.Bool("help3", false, `triangles help`), "kkk"}, strings.Builder{}, true},
		{"Wrong amount of arguments", args{flag.Bool("help4", false, `triangles help`), ""}, strings.Builder{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Task7(tt.args.hf, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Task7() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("Task7() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkNums(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want strings.Builder
	}{
		{"Squares lesser than 10", args{10}, getBuilderWithText("0, 1, 2, 3")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkNums(tt.args.n); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("checkNums() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getBuilderWithText(str string) strings.Builder {
	builder := strings.Builder{}
	builder.WriteString(str)
	return builder
}
