package main 
import "fmt"

func replaceElements(arr []int) []int {
    
	max_ele := arr[len(arr)-1]
	arr[len(arr)-1] = -1

	for iter:= len(arr)-2; iter>=0; iter--{
		temp := arr[iter]
		arr[iter] = max_ele
		if temp > max_ele{
			max_ele = temp
		}
	}

	return arr
}

func main(){

	arr := []int{400,300,500}

	ans := replaceElements(arr)

	fmt.Println(ans)
}