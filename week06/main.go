package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 13              //var i int = 13
	var f float64 = 12.9 //f := 12.9
	c1 := '1'            //49
	c2 := 'あ'            //12354 unicode.org

	fmt.Printf("value i: %d, value f: %f\n", i, f)
	//fmt.Printf("%d + %f = %f\n", i, f, i*f)
	//fmt.Printf("%d * %f = %f\n", i, f, float64(i)*f)
	fmt.Printf("%d * %f = %d\n", i, f, i*int(f))
	fmt.Println(reflect.TypeOf(i), reflect.TypeOf(f))
	fmt.Println(c1, c2)
	fmt.Println(reflect.TypeOf(i), reflect.TypeOf(f), reflect.TypeOf(c1), reflect.TypeOf(c2))
}
