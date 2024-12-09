package hashmap

import (
	"fmt"
	"strings"
)

func PurifyString(word string) string {
	return strings.Join(strings.Split(strings.ToLower(word), " "), "")
}

func CountRepeatChar(word string) *map[rune]int {
	var repchars = make(map[rune]int)
	for _, v := range word {
		repchars[v] += 1
	}
	return &repchars
}

func CountRepeatChar2(word string) *map[rune][]int {
	var repchars = make(map[rune][]int)
	for ix, v := range word {
		repchars[v] = append(repchars[v], ix)
	}
	return &repchars
}

func ReverseRuneInt(runeMap *map[rune]int) map[int][]rune {
	var newMap = make(map[int][]rune)
	for k, v := range *runeMap {
		newMap[v] = append(newMap[v], k)
	}
	return newMap
}

func PrintRuneIntArr(mapsMan *map[rune][]int) {
	for k, v := range *mapsMan {
		fmt.Println(k, "-> ", v)
	}
}

func GetFirstElement(pword string, repCnt int) rune {
	indx := *CountRepeatChar2(pword)
	for ix, v := range pword {
		for rn, ixArr := range indx {
			if len(ixArr) == repCnt {
				if ix == ixArr[0] && rn == v {
					return rn
				}
			}
		}
	}
	return '-'
}

func GetFirstElement2(pword string, repCnt int) rune {
	if indx, OK := ReverseRuneInt(CountRepeatChar(pword))[repCnt]; !OK {
		return '-'
	} else {
		for _, v := range pword {
			for _, rn := range indx {
				if v == rn {
					return v
				} else {
					continue
				}
			}
		}
	}
	return '-'
}
