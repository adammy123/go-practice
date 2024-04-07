package mergeintervals

import (
	"reflect"
	"testing"
)

func Test_mergeIntervals(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test 1",
			args: args{intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}},
			want: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name: "test 2",
			args: args{intervals: [][]int{{1, 4}, {4, 5}}},
			want: [][]int{{1, 5}},
		},
		{
			name: "test 3",
			args: args{intervals: [][]int{{1, 4}, {2, 3}}},
			want: [][]int{{1, 4}},
		},
		{
			name: "test 4",
			args: args{intervals: [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}},
			want: [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
		},
		{
			name: "test 5",
			args: args{intervals: [][]int{{1, 5}, {5, 8}}},
			want: [][]int{{1, 8}},
		},
		{
			name: "test 6",
			args: args{intervals: [][]int{{1, 2}, {1, 8}}},
			want: [][]int{{1, 8}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeIntervals(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortIntervals(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "sort 1",
			args: args{intervals: [][]int{{0, 1}, {1, 2}}},
			want: [][]int{{0, 1}, {1, 2}},
		},
		{
			name: "sort 2",
			args: args{intervals: [][]int{{1, 2}, {0, 1}}},
			want: [][]int{{0, 1}, {1, 2}},
		},
		{
			name: "sort 3",
			args: args{intervals: [][]int{{1, 2}, {1, 3}}},
			want: [][]int{{1, 2}, {1, 3}},
		},
		{
			name: "sort 4",
			args: args{intervals: [][]int{{1, 3}, {1, 2}}},
			want: [][]int{{1, 2}, {1, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortIntervals(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intervalsOverlap(t *testing.T) {
	type args struct {
		i1 []int
		i2 []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "overlap 1",
			args: args{i1: []int{1, 2}, i2: []int{2, 3}},
			want: true,
		},
		{
			name: "overlap 2",
			args: args{i1: []int{2, 3}, i2: []int{1, 2}},
			want: true,
		},
		{
			name: "overlap 3",
			args: args{i1: []int{2, 3}, i2: []int{4, 5}},
			want: false,
		},
		{
			name: "overlap 4",
			args: args{i1: []int{1, 5}, i2: []int{3, 9}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intervalsOverlap(tt.args.i1, tt.args.i2); got != tt.want {
				t.Errorf("intervalsOverlap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_combineTwoIntervals(t *testing.T) {
	type args struct {
		i1 []int
		i2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "combine 1",
			args: args{i1: []int{1, 2}, i2: []int{2, 3}},
			want: []int{1, 3},
		},
		{
			name: "combine 2",
			args: args{i1: []int{2, 3}, i2: []int{1, 2}},
			want: []int{1, 3},
		},
		{
			name: "combine 3",
			args: args{i1: []int{1, 10}, i2: []int{5, 11}},
			want: []int{1, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combineTwoIntervals(tt.args.i1, tt.args.i2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combineTwoIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}
