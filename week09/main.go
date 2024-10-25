package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//fmt.Printf("%f 나누기 %f 은(는) %f입니다.\n", 1.0, 3.0, 1.0/3.0) //바로 출력

	i := 1
	for i <= 10 {
		//fmt.Printf("%3d\n", i)
		fmt.Printf("%4d\n", rand.Intn(1000)+1)
		i++
	}

	//fmt.Printf("%T\n", i)

	result := fmt.Sprintf("%0.2f 나누기 %0.2f 은(는) %0.2f입니다.\n", 1.0, 3.0, 1.0/3.0) //서식 문자열 반환

	fmt.Print(result)
}
