package main

import (
	"fmt"
	"sort"
)

func removeDuplicateLetters(str string) string {

	freqTable := make(map[rune]int)

	unqChar := []rune{}
	for _, char := range str {

		_, ok := freqTable[rune(char)]

		if !ok {
			freqTable[rune(char)] = 1
			unqChar = append(unqChar, rune(char))
		}
	}

	sort.Slice(unqChar, func(i, j int) bool {
		return unqChar[i] < unqChar[j]
	})

	fmt.Println(unqChar)

	ans := ""

	for _, char := range unqChar {
		ans += string(char)
	}

	return ans
}

func main() {

	var str string = "cbacdcbc"

	str = removeDuplicateLetters(str)

	fmt.Println(str)
}
