package main

import (
	"fmt"
	
)

func get_prefix(a string, b string) string{

	iter :=0
	for ; iter<len(a) && iter<len(b); iter++{
		if 	a[iter] != b[iter]{
			return a[:iter]
		} 
	}
	if len(a) < len(b){
		return a[:iter]
	} else{
		return b[:iter]
	}
}


func longestCommonPrefix(strs []string) string {

	ans := strs[0]

	for iter:=1; iter<len(strs); iter++{

		ans = get_prefix(ans, strs[iter])
	}
    
	return ans
}

func main() {

	strs := []string{"abhi", "abhik"}

	ans := longestCommonPrefix(strs)

	fmt.Println(ans)
}