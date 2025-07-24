package stringex

import (
	"fmt"
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

func RevStr(word string) {
	word2 := []rune(word)
	for i := len(word2) - 1; i >= 0; i-- {
		fmt.Printf("%c", word2[i])
	}
	fmt.Println()
}

func RevStr2(word string) string {
	wrdLen := len(word) - 1
	var wordRev = make([]rune, wrdLen+1)
	for i, v := range word {
		wordRev[wrdLen-i] = v
	}
	return string(wordRev)
}

func Palindrome(word string) bool {
	wrdLen := len(word) - 1
	word2 := []rune(MakeCaps(word))
	for i := 0; i <= (wrdLen/2)+1; i++ {
		if word2[i] == word2[wrdLen-i] {
			continue
		} else {
			return false
		}
	}
	return true
}

func CountDuplicate(words string) {
	var letterBits = make([]rune, 26)
	i := 0
	word := []rune(strings.ToLower(words))
	for {
		letterBits[word[i]-97] += 1
		i++
		if i >= len(word) {
			break
		}
	}
	for i, v := range letterBits {
		if v > 1 {
			fmt.Printf("%c:\t%d\n", i+97, v)
		}
	}
}
