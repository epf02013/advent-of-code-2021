package main

import (
	"reflect"
	"testing"
)

func TestParseLine(t *testing.T) {
	type args struct {
		lineString string
	}
	tests := []struct {
		name string
		args args
		want Line
	}{
		{
			name: "Parsing a line",
			args: args{
				lineString: "0,9 -> 5,9",
			},
			want: Line{
				p0: Point{x:0, y: 9},
				p1: Point{x:5, y: 9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseLine(tt.args.lineString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimplifiedSlope(t *testing.T) {
	tests := []struct {
		name string
		line Line
		want Slope
	}{
		{
			name: "Getting simplified slope of line",
			line: Line{
				p0: Point{x:0, y: 0},
				p1: Point{x:6, y: 8},
			},
			want: Slope{x: 3, y: 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.line.simplifiedSlope(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPoints(t *testing.T) {
	tests := []struct {
		name string
		line Line
		want []Point
	}{
		{
			name: "Get Points",
			line: Line{
				p0: Point{x:0, y: 0},
				p1: Point{x:6, y: 8},
			},
			want: []Point{{x: 0, y: 0}, {x: 3, y: 4}, {x: 6, y: 8}},
		},
		{
			name: "Get Points",
			line: Line{
				p0: Point{x:2, y: 2},
				p1: Point{x:2, y: 1},
			},
			want: []Point{{x: 2, y: 2}, {x: 2, y: 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.line.getPoints(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetGCD(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "GCD of 6,8",
			args: args{
				a: 6,
				b: 8,
			},
			want: 2,
		},
		{
			name: "GCD of 15, 10",
			args: args{
				a: 15,
				b: 10,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetGCD(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
