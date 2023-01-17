package main

import (
	"fmt"
)

func main() {
	var limit int = 1000
	var sum int = 0

	for i := 1; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	fmt.Println(sum)
}
