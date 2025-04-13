package mapping

import (
	"fmt"
	"strings"
)

func mapTokens(tokens []string) string {
	var builder strings.Builder
	for _, item := range tokens[:len(tokens)-1] {
		fmt.Fprint(&builder, TransliterationMap[item])
	}

	fmt.Fprint(&builder, getLastTokenMapping(tokens[len(tokens)-1]))

	return builder.String()
}

func getLastTokenMapping(lastToken string) string {
	var res string
	switch lastToken {
	case "a":
		res = TA_MARBUTA
	case "aa":
		res = ALIF_MAQSURA
	default:
		res = TransliterationMap[lastToken]
	}
	return res
}
