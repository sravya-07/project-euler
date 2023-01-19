package main

import (
	"fmt"
	"math"
)

func my_sqrt(a int) int {
	return int(math.Sqrt(float64(a)))
}

func IsPrime(num int) bool {
	if num == 2 {
		return true
	}

	if num%2 == 0 {
		return false
	}

	for i := 3; i < my_sqrt(num); i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	start := 1
	end := 20

	length_of_array := end - start + 1

	// product := 1

	var list_of_nums []int

	for i := start; i <= end; i++ {
		list_of_nums = append(list_of_nums, i)
	}
	fmt.Println(list_of_nums)

	for i := 0; i < length_of_array; i++ {
		for j := i + 1; j < length_of_array; j++ {
			if list_of_nums[j]%list_of_nums[i] == 0 {
				// fmt.Println(list_of_nums[j]/list_of_nums[i], "=", list_of_nums[j], "/", list_of_nums[i])
				list_of_nums[j] = list_of_nums[j] / list_of_nums[i]
			}
		}
	}
	fmt.Println(list_of_nums)

	product := 1
	for i := 0; i < length_of_array; i++ {
		product *= list_of_nums[i]
	}
	fmt.Println(product)
}

// explanation
// 1 2 3 4 5 6 -> 60
// 1 2 3 4 5 6 ->
// 1 2 3 2 5 3
// 1 2 3 2 5 1 ->

// 1 2 3 4 5 6 7 8 ->

// 1 2 3 4 5 6 7 8
// 1 2 3 2 5 6 7 4
// 1 2 3 2 5 2 7 4
// 1 2 3 2 5 1 7 2 840
// 1 2 3 4 5 1 7 2
