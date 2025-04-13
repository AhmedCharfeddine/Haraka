package core

import "testing"

func TestTransliterate(t *testing.T) {
	var testCases = map[string]string{
		"faares": "فارس",
	}

	for latin, want := range testCases {
		var got, _ = Transliterate(latin, MappingStrategy)
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
}

func TestTransliterateShouldNotAcceptShortWord(t *testing.T) {
	var _, err = Transliterate("ab", MappingStrategy)

	if err == nil {
		t.Error("Method should throw error on short words")
	}
}
