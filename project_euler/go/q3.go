package main

func main() {

	var num int64 = 600851475143
	var counter int64 = 3
	var largest_factor int64 = 1
	for num > 1 {
		for num%counter == 0 {
			num = num / counter
			if counter > largest_factor {
				largest_factor = counter
			}
		}

	}
}
