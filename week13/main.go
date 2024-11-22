package main

import (
	"fmt"
)

func main() {
	var empty_slice []int
	//empty_slice = make([]int, 5)
	fmt.Printf("%#v\n", empty_slice)

	arr := [5]float64{3.5, 4.1, 4.5, 3.2, 4.0} //array var
	slice := arr[1:4]                          //slice var

	//slice[1] = 2.71
	arr[2] = 2.71
	slice = append(slice, 4.3, 4.55)
	slice2 := append(slice, 1, 2)

	fmt.Println(slice, slice2)

	slice2[0] = 0

	fmt.Println(slice, slice2)

	fmt.Println(len(slice), slice, arr)
}
