package main

import "fmt"

func isSubsequence(s string, t string) bool {
	idx,jdx := 0, 0
	for idx<len(s) && jdx<len(t){
		if s[idx] == t[jdx]{
			idx++
			jdx++
		} else{
			jdx++
		}
	}

	if idx == len(s){
		return true
	}else{
		return false
	}

}

func main() {
	s, t := "", "ac"

	ans := isSubsequence(s, t)
	fmt.Println(ans)
}