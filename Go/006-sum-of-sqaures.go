package main

import (
	"fmt"
)

func square(num int) int {
	return num * num
}

func main() {
	start := 1
	end := 100
	n := end - start + 1

	sum := (start * n) + (n * (n - 1) / 2)
	// fmt.Println(square(sum))

	// sum of squares
	sum_of_sqaures := 0
	if start != 1 {
		sum_upto_start := (start) * (start + 1) * ((2 * start) + 1) / 6
		sum_upto_end := (end) * (end + 1) * ((2 * end) + 1) / 6

		sum_of_sqaures = sum_upto_end - sum_upto_start + (start * start)
	} else {
		sum_of_sqaures = (n) * (n + 1) * (2*n + 1) / 6
	}
	fmt.Println("square of sum:", square(sum), "sum of sqaures:", sum_of_sqaures, "diff:", square(sum)-sum_of_sqaures)
}
