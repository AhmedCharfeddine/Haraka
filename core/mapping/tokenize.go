package mapping

import (
	"strings"
)

const vowels = "aeiouy"

// word is assumed to have no spaces, and no punctuation
// tokenizes a latin word into tokens for mapping to arabic letters
// prioritizes two-letter keys over one-letter keys
// so that sounds like "sh" are correctly interpreted
// skips vowels, unless they are at the beginning
func Tokenize(word string) []string {
	var tokens []string
	i := 0
	for i < len(word) {
		// Match for two letters
		if i+1 < len(word) {
			twoLetter := word[i : i+2]
			if isValidKey(twoLetter) {
				tokens = append(tokens, twoLetter)
				i += 2
				continue
			}
		}

		// Match for one letter
		oneLetter := string(word[i])
		if isValidKey(oneLetter) {
			if i == 0 || !isVowel(oneLetter) {
				tokens = append(tokens, oneLetter)
			}
		}
		i++
	}
	return tokens
}

func isVowel(char string) bool {
	return strings.Contains(vowels, char)
}
