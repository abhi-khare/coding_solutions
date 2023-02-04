package main

import "fmt"

func isBadVersion(version int) bool {

	if version >= 1 {
		return true
	}

	return false
}

func firstBadVersion(num int) int {

	start, end := 1, num
	if isBadVersion(1) {
		return 1
	}

	for start <= end {
		mid := (start + end) / 2

		fmt.Println(start, end, mid)

		if isBadVersion(mid) && !isBadVersion(mid-1) {
			return mid
		} else if !isBadVersion(mid) {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return start
}

func main() {

	num := 2

	version := firstBadVersion(num)

	fmt.Println(version)

}
