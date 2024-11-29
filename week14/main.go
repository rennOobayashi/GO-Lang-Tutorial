package main

import "fmt"

type visiter struct {
	age   int
	price int
}

// strcuct slice parameter
func calculate_price(visiters []visiter) int {
	var total_price int
	for _, v := range visiters {
		total_price += v.price
	}
	return total_price
}

func main() {
	var numVisiters int
	var vs []visiter

	fmt.Printf("How many visiters: ")
	fmt.Scanln(&numVisiters)

	vs = make([]visiter, numVisiters)

	for i := 0; i < numVisiters; i++ {
		var age int

		fmt.Print("Input age: ")
		fmt.Scanln(&age)

		switch {
		case age < 12:
			vs[i] = visiter{age: age, price: 5000}
		case age < 65:
			vs[i] = visiter{age: age, price: 10000}
		default:
			vs[i] = visiter{age: age, price: 7000}
		}
	}

	fmt.Println("====================")
	fmt.Printf("Total price : %d", calculate_price(vs))
}
