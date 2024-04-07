package vowel

import "testing"

func TestCountNumVowelSubstrings(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"no substrings", "hello", 0},
		{"one substring", "xaeioux", 1},
		{"exactly one substring", "uiaoe", 1},
		{"two substrings", "aaeiou", 2},
		{"too short", "aeio", 0},
		{"complex", "aeiouaeiou", 21},
		{"complex none", "aeioxuaeixou", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountNumVowelSubstrings(tt.input); got != tt.want {
				t.Errorf("CountNumVowelSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
