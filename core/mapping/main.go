package mapping

import (
	"fmt"
	"strings"
)

func TransliterateMapping(latinWord string) string {
	// step 1: convert word to lowercase, remove punctuation
	latinWord = strings.ToLower(latinWord)

	// step 2: break word down by "items" (keys of the map)
	var items = Tokenize(latinWord)

	// step 3: get mappings
	var builder strings.Builder
	for _, item := range items {
		fmt.Fprint(&builder, TransliterationMap[item])
	}

	return builder.String()
}
