package main

import (
	"fmt"
	"os"
	"reflect"
)

func test(i int, strs ...string) {
	fmt.Println(i, strs, reflect.TypeOf(strs))
}

func main() {
	slices := os.Args[1:]

	fmt.Println(slices[1])

	for _, slice := range slices {
		fmt.Println(slice)
	}

	slices = append(slices, "forever", "!")

	fmt.Println(slices, len(slices))

	test(1, "abc")
	test(2, "abc", "dec")
	test(3, "abc", "abc", "abc")
	test(4)
}
