package main

import (
	"fmt"
	"math"
)

func get_factors(number int) []int {
	var result []int

	for i := 2; i < int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			result = append(result, i)
		}
	}
	return result
}

func is_prime(num int) bool {
	for i := 2; i < int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func get_largest_prime_factor(factors []int) int {
	length := len(factors)
	for i := length - 1; i >= 0; i-- {
		if is_prime(factors[i]) {
			return factors[i]
		}
	}
	return 0
}

func main() {
	num := 600851475143
	factors := get_factors(num)
	fmt.Println(factors)
	// fmt.Println(len(factors))
	// for index, value := range factors {
	// 	fmt.Print(index, value, is_prime(value), "\n")
	// }

	larget_prime_factor := get_largest_prime_factor(factors)
	if larget_prime_factor != 0 {
		fmt.Print(larget_prime_factor)
	} else {
		fmt.Print("Has no Prime factors")
	}
}
