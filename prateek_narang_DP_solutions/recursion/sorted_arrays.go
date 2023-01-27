package main 

func checkSorted(arr []int, idx int) bool{

	if idx == len(arr){
		return true
	} else{
		return  (arr[idx]>arr[idx-1]) && checkSorted(arr, idx+1)
	}
}

func main(){

	var arr = []int{4,13}

	status := checkSorted(arr, 1)

	if status{
		println("Array is sorted")
	} else{
		println("Array is not sorted")
	}
}