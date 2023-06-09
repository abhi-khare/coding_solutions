package main

func majorityElement(arr []int) int {
    majorEle := arr[0]
	freq := 1

	for _, ele := range arr {

		if ele == majorEle {
			freq += 1
		} else if ele != majorEle {

			freq -= 1

			if freq == 0 {
				majorEle = ele
				freq = 1
			}
		}
	}

	return majorEle
}