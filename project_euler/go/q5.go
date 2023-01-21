package main

import "fmt"

// compute all the prime numbers less than the target num.
// for each prime number, check the highest exponentiation of number which is smaller or equal to the target num
// final answer is the product of all the exponentiated prime numbers
// Ex for num = 20, it is 2^4 * 3^2 * 5 * 7 * 11 *13 * 17 * 19

func exponentiate(num int64, power int64) int64 {

	var exp_num int64 = 1
	for power >= 1 {
		exp_num = exp_num * int64(num)
		power -= 1
	}

	return exp_num
}

func main() {

	num := 20
	prime_num := []int{2, 3, 5, 7, 11, 13, 17, 19}

	var answer int64 = 1

	for iter := 0; iter < len(prime_num); iter++ {

		power := 1
		exponentiated_prime := prime_num[iter]

		for exponentiated_prime <= num {
			power++
			exponentiated_prime = exponentiated_prime * prime_num[iter]
		}

		answer = answer * exponentiate(int64(prime_num[iter]), int64(power-1))
	}

	fmt.Println("The lowest number that evenly divides all the number from 1-20 is ", answer)

}
