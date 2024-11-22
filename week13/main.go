package main

import (
	"fmt"
)

func main() {
	arr := [5]float64{3.5, 4.1, 4.5, 3.2, 4.0} //array var
	slice := arr[1:4]                          //slice var

	//slice[1] = 2.71
	arr[2] = 2.71
	slice = append(slice, 4.3, 4.55)

	fmt.Println(len(slice), slice, arr)
}
