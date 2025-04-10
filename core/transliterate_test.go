package core

import "testing"

func TestTransliterate(t *testing.T) {
	var testCases = map[string]string{
		"7araka": "حركة",
		"7Araka": "حركة",
		"faares": "فارس",
	}

	for latin, want := range testCases {
		var got = Transliterate(latin, MappingStrategy)
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
}
