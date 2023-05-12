package main

import "fmt"

func isIsomorphic(s string, t string) bool {

	hashmap := make(map[rune]rune)

	if len(s) != len(t){
		return false
	}

	for idx, char := range s{

		val, ok := hashmap[char]
		if !ok{
			hashmap[char] = rune(t[idx])

			for key, value := range hashmap{
				if value == rune(t[idx]) && key != char{
					return false
				}
			}
		} else if val != rune(t[idx]){
			fmt.Print(2, hashmap)
				return false
			}
		}
	

	return true
}

func main() {
	s, t := "paper", "title"

	ans := isIsomorphic(s,t)

	fmt.Print(ans)

}
