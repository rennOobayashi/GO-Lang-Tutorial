package main

import (
	"fmt" //#include <stdio.h>
	"math"
	"strings"
)

func main() {
	//var i int = 55
	var f float32 = 1.1

	i := 55

	fmt.Printf("%d\n", i)
	fmt.Print("1은 ", i, "\n")
	fmt.Println("1은", i)
	fmt.Println("2은", f, math.Ceil(3.49))
	fmt.Println(strings.Title("new title"))
}
