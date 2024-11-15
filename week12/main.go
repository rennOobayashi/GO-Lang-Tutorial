package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr := [5]float64{3.5, 4.1, 4.5, 3.2, 4.0} //array var
	slice := arr[:4]                           //slice var

	fmt.Println(arr, reflect.TypeOf(arr))
	fmt.Println(slice, reflect.TypeOf(slice))
}
