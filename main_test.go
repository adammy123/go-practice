package main

import (
	"testing"
)

func Test_checkGroup(t *testing.T) {
	type args struct {
		winMatches int
		addMatch   bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Group 1",
			args: args{winMatches: 6, addMatch: false},
			want: 1,
		},
		{
			name: "Group 2",
			args: args{winMatches: 5, addMatch: true},
			want: 2,
		},
		{
			name: "Group 3",
			args: args{winMatches: 5, addMatch: false},
			want: 3,
		},
		{
			name: "Group 4",
			args: args{winMatches: 4, addMatch: true},
			want: 4,
		},
		{
			name: "Group 5",
			args: args{winMatches: 4, addMatch: false},
			want: 5,
		},
		{
			name: "Group 6",
			args: args{winMatches: 3, addMatch: true},
			want: 6,
		},
		{
			name: "Group 7",
			args: args{winMatches: 3, addMatch: false},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkGroup(tt.args.winMatches, tt.args.addMatch); got != tt.want {
				t.Errorf("checkGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringSliceContains(t *testing.T) {
	type args struct {
		strSlice []string
		s        string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "strSlice contains s",
			args: args{strSlice: []string{"a", "b"}, s: "a"},
			want: true,
		},
		{
			name: "strSlice contains s",
			args: args{strSlice: []string{"a", "b"}, s: "c"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringSliceContains(tt.args.strSlice, tt.args.s); got != tt.want {
				t.Errorf("stringSliceContains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getGroup(t *testing.T) {
	winningNumbers = []string{"1", "2", "3", "4", "5", "6"}
	additionalNumber = "7"
	type args struct {
		bet []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Group 1",
			args: args{bet: []string{"1", "2", "3", "4", "5", "6"}},
			want: 1,
		},
		{
			name: "Group 2",
			args: args{bet: []string{"1", "2", "3", "4", "5", "7"}},
			want: 2,
		},
		{
			name: "Group 3",
			args: args{bet: []string{"1", "2", "3", "4", "5", "8"}},
			want: 3,
		},
		{
			name: "Group 4",
			args: args{bet: []string{"1", "2", "3", "4", "7", "8"}},
			want: 4,
		},
		{
			name: "Group 5",
			args: args{bet: []string{"1", "2", "3", "4", "8", "9"}},
			want: 5,
		},
		{
			name: "Group 6",
			args: args{bet: []string{"1", "2", "3", "7", "8", "9"}},
			want: 6,
		},
		{
			name: "Group 7",
			args: args{bet: []string{"1", "2", "3", "8", "9", "10"}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGroup(tt.args.bet); got != tt.want {
				t.Errorf("getGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
