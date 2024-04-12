package solution

import (
	"testing"
)

func Test_countTriplets(t *testing.T) {
	type args struct {
		arr []int64
		r   int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "test 1",
			args: args{arr: []int64{1, 2, 2, 4}, r: 2},
			want: 2,
		},
		{
			name: "test 2",
			args: args{arr: []int64{1, 3, 9, 9, 27, 81}, r: 3},
			want: 6,
		},
		{
			name: "test 3",
			args: args{arr: []int64{1, 5, 5, 25, 125}, r: 5},
			want: 4,
		},
		{
			name: "test 4",
			args: args{arr: []int64{2, 5, 5, 25, 125}, r: 5},
			want: 2,
		},
		{
			name: "test 5",
			args: args{arr: []int64{125, 5, 5, 25, 1}, r: 5},
			want: 0,
		},
		{
			name: "test 6",
			args: args{arr: []int64{1, 1, 1, 2, 2, 2, 2}, r: 1},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countTriplets(tt.args.arr, tt.args.r); got != tt.want {
				t.Errorf("countTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLowestDenominator(t *testing.T) {
	type args struct {
		val        int64
		r          int64
		multiplier int
	}
	tests := []struct {
		name  string
		args  args
		want  int64
		want1 int
	}{
		{
			name:  "test 1",
			args:  args{10, 2, 1},
			want:  5,
			want1: 2,
		},
		{
			name:  "test 2",
			args:  args{20, 2, 1},
			want:  5,
			want1: 3,
		},
		{
			name:  "test 3",
			args:  args{20, 8, 1},
			want:  20,
			want1: 1,
		},
		{
			name:  "test 4",
			args:  args{16, 2, 1},
			want:  1,
			want1: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getLowestDenominator(tt.args.val, tt.args.r, tt.args.multiplier)
			if got != tt.want {
				t.Errorf("getLowestDenominator() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getLowestDenominator() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_binomial(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "100C3",
			args: args{n: 100, k: 3},
			want: 161700,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binomial(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("binomial() = %v, want %v", got, tt.want)
			}
		})
	}
}
