package solution

import (
	"testing"
)

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test 1",
			args: args{s: "aabbcd"},
			want: "NO",
		},
		{
			name: "test 2",
			args: args{s: "aabbccddeefghi"},
			want: "NO",
		},
		{
			name: "test 3",
			args: args{s: "abcdefghhgfedecba"},
			want: "YES",
		},
		{
			name: "test 4",
			args: args{s: "a"},
			want: "YES",
		},
		{
			name: "test 5",
			args: args{s: "aaaaa"},
			want: "YES",
		},
		{
			name: "test 6",
			args: args{s: "aaabbbcccdddee"},
			want: "NO",
		},
		{
			name: "test 7",
			args: args{s: "aaabbbcccdddeeee"},
			want: "YES",
		},
		{
			name: "test 8",
			args: args{s: "aaaabbcc"},
			want: "NO",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
