package main

import (
	"fmt"
	"strconv"
	"strings"
)

func decodeString(str string) string {

	ans := ""
	tokens := strings.Split(str, "]")

	for _, words := range tokens {
		split_word := strings.Split(words, "[")

		count, _ := strconv.Atoi(split_word[0])

		for jter := 0; jter < count; jter++ {
			ans += split_word[1]
		}
	}

	return ans
}

func main() {

	str := "3[a]2[bc]"

	str = decodeString(str)

	fmt.Println(str)
}
