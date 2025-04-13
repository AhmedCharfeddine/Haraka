package core

import (
	"fmt"

	"github.com/AhmedCharfeddine/Haraka/core/mapping"
)

const MappingStrategy = "mapping"

func Transliterate(latinWord string, strategy string) (arabicWord string, e error) {
	if len(latinWord) < 3 {
		return latinWord, fmt.Errorf("word has to be at least three characters long")
	}
	if strategy == MappingStrategy {
		return mapping.TransliterateMapping(latinWord), nil
	}
	return latinWord, fmt.Errorf("unknown transliteration strategy")
}
