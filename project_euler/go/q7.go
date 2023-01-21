package main

import "fmt"

// using prime number theorem, we can say get a rough upper bound on
// the limit of numbers that will have 10001 prime numbers. We can start with
// the limit as 600000
// then we will apply sieve of eratosthenes to find the nth prime number.

func main() {

	var non_prime_set = new([600000]int)

	non_prime_set[0] = 1
	non_prime_set[1] = 1

	prime_count := 0

	for iter := 2; iter < 600000; iter++ {

		if non_prime_set[iter] == 0 {

			prime_count++

			if prime_count == 10001 {

				fmt.Println("10001th prime is ", iter)

				return
			}

			for jter := iter * iter; jter < 600000; jter = jter + iter {
				non_prime_set[jter] = 1
			}
		}

	}
}
