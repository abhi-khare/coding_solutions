package main 

import "fmt"

func maxNumberOfBalloons(text string) int {
    
	hashMap := make(map[rune]int, 5)
	for _, ele := range "balon"{
		hashMap[ele] = 0
	}

	for _, ele := range text{
		if val, ok := hashMap[ele]; ok{
			hashMap[ele] = val +1
		}
	}

	minCount := 100000000
	for _, ele := range []int{'b', 'a', 'n', 'o', 'l'}{
		count := hashMap[rune(ele)]
		if ele == 'l' || ele =='o'{
			if int(count/2) < minCount{
				minCount = int(count/2)
			}
		} else{
			if count < minCount{
				minCount = count
			}

		} 

	}

	return minCount



	
}


func main(){

	text := "b"

	ans := maxNumberOfBalloons(text)

	fmt.Println(ans)
}