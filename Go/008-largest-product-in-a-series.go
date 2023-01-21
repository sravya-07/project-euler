package main

import (
	"fmt"
	"os"
	"reflect"
	// "runtime"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := os.ReadFile("inputs/008.txt")
	check(err)
	fmt.Println(data)
	fmt.Println(reflect.TypeOf(data))
}
