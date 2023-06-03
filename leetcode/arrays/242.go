package main

import "fmt"

func isAnagram(s string, t string) bool {

	if len(s) != len(t){
		return false
	}

	freqMap := make(map[rune]int)

	for iter:=0; iter<len(s); iter++{

		s_char := rune(s[iter])
		t_char := rune(t[iter])
		

		if _, ok := freqMap[s_char]; ok{
			freqMap[s_char] +=1
		} else{
			freqMap[s_char] = 1
		}

		if _, ok := freqMap[t_char]; ok{
			freqMap[t_char] -=1
		} else{
			freqMap[t_char] = -1
		}
	}

	for _, val := range freqMap{
		if val !=0{
			return false
		}
	}
	return true
}

func main() {

	s := "aaa"
	t := "aab"

	ans := isAnagram(s, t)
	fmt.Println(ans)
}