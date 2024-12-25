package stringex

func MakeCaps(word string) string {
	var newWord = make([]rune, len(word))
	for i, v := range word {
		newWord[i] = v - 32
	}
	word = string(newWord)
	// return string(newWord)
	return word

}
