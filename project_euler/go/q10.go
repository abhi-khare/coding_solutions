package main

import "fmt"

func main() {

	var non_prime_set = new([2000000]int)

	non_prime_set[0] = 1
	non_prime_set[1] = 1

	var prime_sum int64 = 0
	for iter := 2; iter < 2000000; iter++ {

		if non_prime_set[iter] == 0 {

			prime_sum += int64(iter)

			for jter := iter * iter; jter < 2000000; jter = jter + iter {
				non_prime_set[jter] = 1
			}

		}
	}

	fmt.Println("sum of prime number is ", prime_sum)
}
