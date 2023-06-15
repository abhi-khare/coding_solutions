package main

import (
	"fmt"
	"strings"
)


func wordPattern(pattern string, str string) bool {

	tokens := strings.Fields(str)
	
	if len(pattern) != len(tokens){
		return false
	}

	forwardMap := make(map[rune]string)
	reverseMap := make(map[string]rune)

	for iter:=0; iter<len(pattern); iter++{
		
		char := rune(pattern[iter])
		word := tokens[iter]

		if val, ok := forwardMap[char]; ok{
			if val != word{
				return false
			} 
		} else{
			forwardMap[char] = word

			if val, ok := reverseMap[word]; ok{
				if val != char {
					return false
				}
			} else{
				reverseMap[word] = char
			}
			
		}
	}

	return true
       
}

func main(){

pattern := "aabb"
s := "dog dog cat dog"

ans := wordPattern(pattern, s)

fmt.Println(ans)


}

