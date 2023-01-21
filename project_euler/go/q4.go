package main

func isPlaindrome(num int) bool {
	if num%10 == num/100000 && (num/10)%10 == (num/10000)%10 && (num/100)%10 == (num/1000)%10 {
		return true
	} else {
		return false
	}
}

func main() {

	var largestPalindrome int = 0

	for i := 999; i >= 101; i = i - 1 {

		var b int
		var diff int

		if i%11 == 0 {
			b = 999
			diff = 1
		} else {
			b = 990
			diff = 11
		}

		for j := b; j >= i; j = j - diff {
			if i*j <= largestPalindrome {
				break
			}
			if isPlaindrome(i * j) {
				largestPalindrome = i * j
			}
		}

	}
}
