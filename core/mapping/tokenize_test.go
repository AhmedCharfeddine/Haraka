package mapping

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"khaaled", []string{"kh", "aa", "l", "d"}},
		{"a7med", []string{"a", "7", "m", "d"}},
		{"haytham", []string{"h", "y", "th", "m"}},
		{"faares", []string{"f", "aa", "r", "s"}}, // "a" is removed (mid-vowel)
		{"7araka", []string{"7", "r", "k"}},
		{"shams", []string{"sh", "m", "s"}},
	}

	for _, test := range tests {
		result := Tokenize(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("tokenize(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}
