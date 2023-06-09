package main

import "fmt"

func isIsomorphic(s string, t string) bool {

	hashMap := make(map[rune]rune)

	if len(s) != len(t){
		return false
	}

	for idx, ele := range s{

		if val, ok := hashMap[ele]; ok{

			if val != rune(t[idx]){
				return false
			}
		} else {

			for _, value := range hashMap {
				if value == rune(t[idx]){
					return false
				}
			}

			hashMap[ele] = rune(t[idx])

		}
	}
	return true
}

func main() {
	s, t := "ff", "bg"

	ans := isIsomorphic(s, t)

	fmt.Println(ans)
}