package main 
import "fmt"

func twoSum(nums []int, target int) []int {
    
	hashMap := make(map[int]int)

	for idx, ele := range nums{

		residual := target - ele

		if val, ok := hashMap[residual]; ok{
			return []int{idx, val}
		} else{
			hashMap[ele] = idx
		}
	}

	return []int{0,0}
}


func main(){

	nums := []int{3,-3 ,0}
	target := 0

	ans := twoSum(nums, target)
	fmt.Println(ans)
}