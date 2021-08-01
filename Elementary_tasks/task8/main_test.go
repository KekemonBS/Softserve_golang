package main

import (
	"flag"
	"testing"
)

func TestTask7(t *testing.T) {
	type args struct {
		hf  *bool
		arg [2]string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Return basic range from 5 to 100 :", args{flag.Bool("help1", false, `triangles help`), [2]string{"5", "100"}}, "5, 8, 13, 21, 34, 55, 89", false},
		{"Display help", args{flag.Bool("help2", true, `triangles help`), [2]string{}}, "", false},
		{"Wrong amouunt of arguments eror", args{flag.Bool("help3", false, `triangles help`), [2]string{}}, "", true},
		{"Wrong first argument", args{flag.Bool("help4", false, `triangles help`), [2]string{"5k", "100"}}, "", true},
		{"Wrong second argument", args{flag.Bool("help5", false, `triangles help`), [2]string{"5", "100k"}}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Task7(tt.args.hf, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Task7() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Task7() = %v, want %v", got, tt.want)
			}
		})
	}
}
