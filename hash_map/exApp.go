package hashmap

import (
	"fmt"
	"strings"
)

func purifyString(word string) string {
	return strings.Join(strings.Split(strings.ToLower(word), " "), "")
}

func CountRepeatChar(word string) *map[rune]int {
	var repchars = make(map[rune]int)
	pword := purifyString(word)
	for _, v := range pword {
		repchars[v] += 1
	}
	return &repchars
}

func PrintRuneMap(runeMap *map[rune]int) {
	for k, v := range *runeMap {
		fmt.Printf("%c -> %d\n", k, v)
	}
}

// hashmap.PrintRuneMap(hashmap.CountRepeatChar("               sjldkbfkjsf pewytopwqytopqyt ,nbnm cdvdxv"))
// Get the first non repetative letter
func getFirstNonrepeatElement(runeMap map[rune]int) []rune {
	var repeat []rune
	for k, _ := range runeMap {
		if runeMap[k] == 1 {
			repeat = append(repeat, k)
		}
	}
	return repeat
}
