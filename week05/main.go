package main

import (
	"fmt" //#include <stdio.h>
	"math"
	"reflect"
	"strings"
)

func main() {
	//명확하게 정의할 때 사용
	//var i int = 55
	//var i int
	//i = 55
	//var f float32 = 1.1

	//i := 55
	i := "55"
	f := 1.1 //64로 됨

	fmt.Printf("%s\n", i)
	fmt.Print("1은 ", i, "\n")
	fmt.Println("1은", i)
	fmt.Println("2은", f, math.Ceil(3.49))
	fmt.Println(strings.Title("new title"))
	fmt.Println(reflect.TypeOf(i), reflect.TypeOf(f))
}
