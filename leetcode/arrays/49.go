package main

import "fmt"
import "strconv"

func getHashString(word string) string{
	freqArray := make([]int, 26)

	for _, char := range(word){
		freqArray[rune(char)-'a'] +=1
	}

	hashStr := ""
	for _, count := range(freqArray){
		hashStr += strconv.Itoa(count) + "#"
	}

	return hashStr

}

func groupAnagrams(strs []string) [][]string {

	hashMap := make(map[string][]string)
	ans := [][]string{}

	for _, word := range strs{
		hashString := getHashString(word)
		
		if _, ok := hashMap[hashString]; ok {
			hashMap[hashString] = append(hashMap[hashString], word)
		} else{
			hashMap[hashString] = []string{word}
		}
	}

	for _, value := range hashMap{
		ans = append(ans, value)
	}
	return ans
}

func main() {
	strs := []string{"","abh", "hab", "abhi"}

	strs_group := groupAnagrams(strs)

	fmt.Println(strs_group)

}