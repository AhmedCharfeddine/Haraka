package mapping

import (
	"strings"
)

func TransliterateMapping(latinWord string) string {
	latinWord = strings.ToLower(latinWord)
	var tokens = Tokenize(latinWord)
	var result = mapTokens(tokens)
	return result
}
