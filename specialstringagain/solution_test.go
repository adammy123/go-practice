package solution

import (
	"testing"
)

func Test_substrCount(t *testing.T) {
	type args struct {
		n int32
		s string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "test 1",
			args: args{n: 5, s: "asasd"},
			want: 7,
		},
		{
			name: "test 2",
			args: args{n: 7, s: "abcbaba"},
			want: 10,
		},
		{
			name: "test 3",
			args: args{n: 4, s: "aaaa"},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := substrCount(tt.args.n, tt.args.s); got != tt.want {
				t.Errorf("substrCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_palindromeNum(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "palindrome 1",
			args: args{1},
			want: 1,
		},
		{
			name: "palindrome 3",
			args: args{3},
			want: 3,
		},
		{
			name: "palindrome 5",
			args: args{5},
			want: 6,
		},
		{
			name: "palindrome 7",
			args: args{7},
			want: 10,
		},
		{
			name: "palindrome 9",
			args: args{9},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := palindromeNum(tt.args.n); got != tt.want {
				t.Errorf("palindromeNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trianglularNum(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "triangularNum 1",
			args: args{1},
			want: 1,
		},
		{
			name: "triangularNum 2",
			args: args{2},
			want: 3,
		},
		{
			name: "triangularNum 3",
			args: args{3},
			want: 6,
		},
		{
			name: "triangularNum 4",
			args: args{4},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trianglularNum(tt.args.n); got != tt.want {
				t.Errorf("trianglularNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
