package main

import "fmt"

func sieve_prime() []int {

	var non_prime_set = new([600000]int)

	non_prime_set[0] = 1
	non_prime_set[1] = 1

	var prime_list []int = make([]int, 0)

	for iter := 2; iter < 600000; iter++ {

		if non_prime_set[iter] == 0 {

			// store the prime
			prime_list = append(prime_list, iter)

			for jter := iter * iter; jter < 600000; jter = jter + iter {
				non_prime_set[jter] = 1

			}
		}
	}

	return prime_list
}

func compute_divisor(num int, primes []int) int {

	num_div := 1

	for iter := 0; iter < len(primes); iter++ {

		factor_count := 0
		for num%primes[iter] == 0 {
			factor_count++
			num = num / primes[iter]
		}
		num_div = num_div * (factor_count + 1)

		if num == 1 {
			break
		}

	}
	return num_div
}

func main() {

	prime := sieve_prime()

	for iter := (30); ; iter++ {

		tri_num := iter * (iter + 1) / 2

		num_div := compute_divisor(tri_num, prime)

		if num_div > 500 {
			fmt.Println("Triangular number is ", tri_num)
			break
		}
	}
}
