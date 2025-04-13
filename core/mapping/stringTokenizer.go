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
	length := len(word)
	for i < length {
		// Match for two letters
		if i+1 < length {
			twoLetter := word[i : i+2]
			if isValidKey(twoLetter) {
				tokens = append(tokens, twoLetter)
				i += 2
				continue
			}
		}

		// Match for one letter
		oneLetter := string(word[i])
		var lastLetter string
		if i != 0 {
			lastLetter = string(word[i-1])
		}
		if isValidKey(oneLetter) {
			if !isSkippableVowel(i, oneLetter, lastLetter, length) {
				tokens = append(tokens, oneLetter)
			}
		}
		i++
	}
	return tokens
}

// We should generally skip mapping vowels in words.
// With two exceptions:
// 1. If the vowel is at the start; not to skip 'a' in 'achref'
// 2. If the vowel is preceeded by another vowel, not to skip 'ouya' in '5ouya'
// 3. If the vowel is at the end; to interpret ta marbuta
func isSkippableVowel(charPos int, char string, lastChar string, length int) bool {
	return strings.Contains(vowels, char) &&
		charPos != 0 &&
		!strings.Contains(vowels, lastChar) &&
		charPos < length-1
}
