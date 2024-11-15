package main

import (
	"fmt"
	"log"
	"week11/keyboard"
	//"github.com/headfirstgo/keyboard"
)

func isPrime(number int) bool {
	//1 is not prime number
	if number <= 1 {
		return false
	} else if number == 2 {
		return true
	} else if number%2 == 0 {
		return false
	} else {
		j := 3
		for j*j <= number {
			if number%j == 0 {
				return false
			}

			j += 2
		}
	}
	return true
}

func main() {

	fmt.Print("input start number: ")
	n1, err1 := keyboard.GetInteger()
	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Print("input end number: ")
	n2, err2 := keyboard.GetInteger()
	if err2 != nil {
		log.Fatal(err2)
	}

	for i := n1; i <= n2; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
}
