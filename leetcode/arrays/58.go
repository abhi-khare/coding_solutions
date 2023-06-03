package main

import "fmt"

func lengthOfLastWord(s string) int {
    ans, flag := 0, 0

	for idx := len(s)-1;idx>=0;idx--{
		if rune(s[idx]) != ' '{
			flag = 1
		} 
		
		if flag ==1 && rune(s[idx])==' '{
			return ans
		} else if flag ==1{
			ans++
		}

	}
	return ans
}

func main() {

	s := "H HH abc "

	ans := lengthOfLastWord(s)

	fmt.Println(ans)
}