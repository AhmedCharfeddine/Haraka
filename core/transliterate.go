package core

import (
	"github.com/AhmedCharfeddine/Haraka/core/mapping"
)

const MappingStrategy = "mapping"

func Transliterate(latinWord string, strategy string) string {
	if strategy == MappingStrategy {
		return mapping.TransliterateMapping(latinWord)
	}
	return latinWord
}
