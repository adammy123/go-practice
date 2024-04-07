package decrypt

import (
	"testing"
)

func TestDescrypt(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "dnotq to crime",
			args: args{word: "dnotq"},
			want: "crime",
		},
		{
			name: "flgxswdliefy to encyclopedia",
			args: args{word: "flgxswdliefy"},
			want: "encyclopedia",
		},
		{
			name: "a to z",
			args: args{word: "a"},
			want: "z",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decrypt(tt.args.word); got != tt.want {
				t.Errorf("Decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertCharToAscii(t *testing.T) {
	type args struct {
		char rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "c to 99",
			args: args{rune('c')},
			want: 99,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertCharToAscii(tt.args.char); got != tt.want {
				t.Errorf("convertCharToAscii() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertAsciiToString(t *testing.T) {
	type args struct {
		ascii int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "99 to c",
			args: args{99},
			want: "c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertAsciiToString(tt.args.ascii); got != tt.want {
				t.Errorf("convertAsciiToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
