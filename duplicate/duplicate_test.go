package duplicate

import "testing"

func TestRemoveDuplicate(t *testing.T) {
	type args struct {
		input string
		k     int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "assassinate to inate with k=2",
			args: args{"assassinate", 2},
			want: "inate",
		},
		{
			name: "mississippi to m with k=2",
			args: args{"mississippi", 2},
			want: "m",
		},
		{
			name: "mississippi to mississippi with k=3",
			args: args{"mississippi", 3},
			want: "mississippi",
		},
		{
			name: "aa to nothing with k=2",
			args: args{"aa", 2},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDuplicate(tt.args.input, tt.args.k); got != tt.want {
				t.Errorf("RemoveDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
