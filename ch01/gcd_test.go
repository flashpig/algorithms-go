package main

import "testing"

func Test_gcd(t *testing.T) {
	type args struct {
		p int
		q int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"gcd", args{54, 63}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcd(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gcd2(t *testing.T) {
	type args struct {
		p int
		q int
	}
	tests := []struct {
		name  string
		args  args
		wantR int
	}{
		// TODO: Add test cases.
		{"gcd", args{54, 63}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := gcd2(tt.args.p, tt.args.q); gotR != tt.wantR {
				t.Errorf("gcd2(%d, %d) = %v, want %v", tt.args.p, tt.args.q, gotR, tt.wantR)
			}
		})
	}
}
