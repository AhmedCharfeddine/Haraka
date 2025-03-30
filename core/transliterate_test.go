package core

import "testing"

func TestTransliterate(t *testing.T) {
	var testCases = map[string]string{
		"haraka": "حركة",
		"Haraka": "حركة",
		"fares":  "فارس",
	}

	for latin, want := range testCases {
		var got = Transliterate(latin)
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
}
