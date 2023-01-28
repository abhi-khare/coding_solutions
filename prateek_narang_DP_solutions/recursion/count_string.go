package main

import "fmt"

func getPossibleWays(str string, idx int, len int) int {

	if idx == len-1{
		return 1
	} else if str == ""{
		c1 := getPossibleWays(str + "0", idx+1, len)
		c2 := getPossibleWays(str + "1", idx+1, len)

		return c1 + c2
	} else{
		if str[idx] == '0'{
			c1 := getPossibleWays(str + "0", idx+1, len)
			c2 := getPossibleWays(str + "1", idx+1, len)

			return c1 + c2
		} else {
			c1 := getPossibleWays(str + "0", idx+1, len)
			return c1
		}
	}
}

func main() {

	var len int
	fmt.Scan(&len)

	str := ""

	count := getPossibleWays(str, -1, len)

	fmt.Print(count)
}