package main

func main() {

	nums := []int{2, 11, 7, 15}
	hashMap := make(map[int]int)

	target := 9

	for iter, ele := range nums {
		idx, ok := hashMap[target-ele]
		if ok {
			print(iter, idx)
		} else {
			hashMap[ele] = iter
		}
	}

}
