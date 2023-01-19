package main

import (
	"fmt"
	"math"
)

func my_pow(a int, n int) int {
	return int(math.Pow(float64(a), float64(n)))
}

func IsPalindrome(num int) bool {
	reverse, temp := 0, num

	// if temp%10 == 0 {
	// 	reverse = 10
	// }
	// fmt.Println("num", num)
	for temp > 0 {
		digit := temp % 10

		reverse = (reverse * 10) + digit
		temp = temp / 10
		// fmt.Println("temp", temp, "rev", reverse)
	}

	// fmt.Println("-----------------")

	// fmt.Println("num:", num, "reverse:", reverse)
	if num == reverse {
		return true
	} else {
		return false
	}
}

func main() {

	start := 100
	end := 999

	// var result []int
	var max int = 0
	// var flag bool = false

	for i := end; i >= start; i-- {
		for j := end; j >= start; j-- {
			if IsPalindrome(i * j) {
				fmt.Println("num 1", i, "num2", j, i*j)
				if i*j > max {
					max = i * j
				}
				// result = append(result, i*j)
			}
		}
	}
	print(max)
}
