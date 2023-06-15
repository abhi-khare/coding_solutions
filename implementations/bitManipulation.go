package main

import "fmt"
import "math"

func main() {
	//arr := make([]int, 125)
	a := 10%3
	out := math.Pow(2, float64(a))
	fmt.Println(out)
}