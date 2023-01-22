package main

import (
	"strings"
)

func checkPattern(pattern string, str string) bool {

	strMap := make(map[string]string)
	revStrMap := make(map[string]string)

	tokens := strings.Fields(str)

	if len(pattern) != len(tokens){
		return false
	}

	for iter:=0; iter< len(pattern); iter++{

		char := pattern[iter:iter+1]
		val, ok := strMap[char]

		if ok == false{
			strMap[char] = tokens[iter]
		} else{
			if val != tokens[iter]{
				return false
			}
		}
	}

	for iter:=0; iter< len(tokens); iter++{

		char := pattern[iter:iter+1]
		val, ok := revStrMap[tokens[iter]]

		if ok == false{
			revStrMap[tokens[iter]] = char
		} else{
			if val != char{
				return false
			}
		}
	}
	
	return true
}
