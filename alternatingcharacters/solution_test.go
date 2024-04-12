package solution

import "testing"

func Test_alternatingCharacters(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "test 1",
			args: args{s: "AAAA"},
			want: 3,
		},
		{
			name: "test 2",
			args: args{s: "BBBBB"},
			want: 4,
		},
		{
			name: "test 3",
			args: args{s: "ABABABAB"},
			want: 0,
		},
		{
			name: "test 4",
			args: args{s: "BABABABA"},
			want: 0,
		},
		{
			name: "test 5",
			args: args{s: "AAABBB"},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alternatingCharacters(tt.args.s); got != tt.want {
				t.Errorf("alternatingCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
