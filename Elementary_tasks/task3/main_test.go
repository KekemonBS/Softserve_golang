package main

import (
	"errors"
	"flag"
	"reflect"
	"testing"
)

func Test_triangle_area(t *testing.T) {
	type fields struct {
		name string
		s1   float64
		s2   float64
		s3   float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
		{"Triangle (1, 1, 1)", fields{"(1, 1, 1)", 1, 1, 1}, 0.43},
		{"Triangle (3, 4, 5)", fields{"(3, 4, 5)", 3, 4, 5}, 6.0},
		{"Triangle (6, 8, 10)", fields{"(6, 8, 10)", 6, 8, 10}, 24.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := triangle{
				name: tt.fields.name,
				s1:   tt.fields.s1,
				s2:   tt.fields.s2,
				s3:   tt.fields.s3,
			}
			if got := tr.area(); got != tt.want {
				t.Errorf("triangle.area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortShapes(t *testing.T) {
	type args struct {
		sh shapeList
	}
	tests := []struct {
		name string
		args args
		want shapeList
	}{
		{"No sorting needed", args{shapeList{triangle{"First", 3, 4, 5}}},
			shapeList{triangle{"First", 3, 4, 5}}},
		{"Sort two shapes", args{shapeList{triangle{"Second", 3, 4, 5}, triangle{"First", 6, 8, 10}}},
			shapeList{triangle{"First", 6, 8, 10}, triangle{"Second", 3, 4, 5}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortShapes(tt.args.sh); got[0] != tt.want[0] {
				t.Errorf("compareTwoLetters() = %v, want %v", got, tt.want)
			}
		})
	}

}

func Test_convertShapes(t *testing.T) {
	type args struct {
		ss string
	}

	type want struct {
		shl shapeList
		err error
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{"One shape", args{"First,   3, 4,  5\n"}, want{shapeList{triangle{"First", 3, 4, 5}}, nil}},
		{"Several shapes", args{"Second,   3, 4,  5\nFirst, 6, 8, 10\n"},
			want{shapeList{triangle{"Second", 3, 4, 5}, triangle{"First", 6, 8, 10}}, nil}},
		{"Wrong first side", args{"First,   3d, 4,  5\n"}, want{shapeList{}, errors.New("wrong first side entered")}},
		{"Wrong second side", args{"First,   3, 4d,  5\n"}, want{shapeList{}, errors.New("wrong second side entered")}},
		{"Wrong third side", args{"First,   3, 4,  5d\n"}, want{shapeList{}, errors.New("wrong third side entered")}},
		{"Wrong sides", args{"First,   1, 4,  5\n"}, want{shapeList{}, errors.New("the entered sides cannot be combined into a triangle")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := convertShapes(tt.args.ss); reflect.DeepEqual(got, tt) && err.Error() == tt.want.err.Error() {
				t.Errorf("compareTwoLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_printSortedShapes(t *testing.T) {
	type args struct {
		sh   shapeList
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Print 2 shapes",
			args{shapeList{triangle{"T2", 6, 8, 10}, triangle{"T1", 3, 4, 5}}, "Triangle"},
			"\n============= Triangle list: ===============\n1. [Triangle T2]: 24.00 cm^2\n2. [Triangle T1]: 6.00 cm^2\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := printSortedShapes(tt.args.sh, tt.args.name)
			if got != tt.want {
				t.Errorf("printSortedShapes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTask3(t *testing.T) {
	type args struct {
		hf       *bool
		proposed string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"No failure, help flag", args{flag.Bool("help1", true, `triangles help`), ""}, false},
		{"No failure, no help flag", args{flag.Bool("help2", false, `triangles help`), "Name, 3, 4, 5\n"}, false},
		{"Failure, no help flag", args{flag.Bool("help3", false, `triangles help`), "Name, 1, 2, 10\n"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Task3(tt.args.hf, tt.args.proposed); (err != nil) != tt.wantErr {
				t.Errorf("Task3() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
