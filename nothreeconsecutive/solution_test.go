package solution

import (
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		A int
		B int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test 1",
			args: args{A: 5, B: 3},
		},
		{
			name: "test 2",
			args: args{A: 3, B: 3},
		},
		{
			name: "test 3",
			args: args{A: 1, B: 4},
		},
		{
			name: "test 4",
			args: args{A: 100, B: 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.A, tt.args.B); !isStringValid(got, tt.args.A, tt.args.B) {
				t.Errorf("Solution(%d, %d) = %v, is not a valid solution!", tt.args.A, tt.args.B, got)
			}
		})
	}
}

func Test_isStringValid(t *testing.T) {
	type args struct {
		res string
		A   int
		B   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test 1a",
			args: args{res: "aabaabab", A: 5, B: 3},
			want: true,
		},
		{
			name: "test 1b",
			args: args{res: "abaabbaa", A: 5, B: 3},
			want: true,
		},
		{
			name: "test 1c",
			args: args{res: "ababbaa", A: 5, B: 3},
			want: false,
		},
		{
			name: "test 1d",
			args: args{res: "ababbaaa", A: 5, B: 3},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStringValid(tt.args.res, tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("isStringValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
