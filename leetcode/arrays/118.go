package main

import "fmt"

func generate(numRows int) [][]int {
	
	ans := make([][]int, numRows)

	if numRows == 1{
		ans[0] = append(ans[0], 1)
	} else if numRows == 2{
		ans[0] = append(ans[0], 1)
		ans[1] =  []int{1,1}
	} else{
		ans[0] = append(ans[0], 1)
		ans[1] =  []int{1,1}
		for iter:=2; iter<numRows;iter++{
			row:= []int{1}
			for jter:=0; jter<iter-1;jter++{
				row = append(row, ans[iter-1][jter] + ans[iter-1][jter+1])
			}
			row = append(row, 1)

			ans[iter] = row
		}

	}

	return ans
}

func main() {

	numRows := 5

	ans := generate(numRows)

	fmt.Println(ans)
}