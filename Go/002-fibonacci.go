package main

import (
	"fmt"
)

func main() {
	limit := 4000000
	f1 := 1
	f2 := 2
	sum := 0

	for f2 < limit {
		if f2%2 == 0 {
			// fmt.Print(f2, " ")
			sum += f2
		}
		temp := f1
		f1 = f2
		f2 = temp + f2
	}

	fmt.Println(sum)
}
