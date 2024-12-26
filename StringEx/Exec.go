package stringex

import (
	"strings"
)

func MakeCaps(word string) string {
	var newWord = make([]rune, len(word))
	for i, v := range word {
		newWord[i] = v - 32
	}
	word = string(newWord)
	return word
}

func CountVouls(word string) (int, int, int) {
	var consonants, vouls, words int
	var prv_spc rune
	word = strings.TrimSpace(word)
	for i, v := range word {
		switch rune(v) {
		case 'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U':
			vouls += 1
		case ' ':
			prv_spc = rune(word[i-1])
			if prv_spc != ' ' {
				words += 1
			}
		default:
			consonants += 1
		}
	}
	return vouls, consonants, words + 1
}

func ValidChar(secword string) bool {
	secword = strings.TrimSpace(secword)
	for _, i := range secword {
		if i >= 48 && i <= 57 || i >= 65 && i <= 90 || i >= 97 && i <= 122 {
			continue
		} else {
			return false
		}
	}
	return true
}
