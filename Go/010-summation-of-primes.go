package main

import (
	"fmt"
)

func GetRange(min int, max int) []int {
	var result []int
	for i := min; i <= max; i++ {
		result = append(result, i)
	}
	return result
}

func main() {
	end := 2000000
	sieve := GetRange(1, end)

	for i := 1; i < end; i++ {
		for j := i + 1; j < end; j++ {
			if sieve[j]%sieve[i] == 0 {
				sieve[j] = sieve[j] / sieve[i]
			}
		}
	}
	fmt.Println(sieve)
}
