package distinctstring

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		P string
		Q string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{P: "ca", Q: "ab"},
			want: 1,
		},
		{
			name: "test 2",
			args: args{P: "abc", Q: "bcd"},
			want: 2,
		},
		{
			name: "test 3",
			args: args{P: "axxz", Q: "yzwy"},
			want: 2,
		},
		{
			name: "test 4",
			args: args{P: "bacad", Q: "abada"},
			want: 1,
		},
		{
			name: "test 5",
			args: args{P: "amz", Q: "amz"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.P, tt.args.Q); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
