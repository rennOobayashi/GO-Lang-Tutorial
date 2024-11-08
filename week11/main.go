package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	in := bufio.NewReader(os.Stdin)
	a, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	a = strings.TrimSpace(a)
	n1, _ := strconv.Atoi(a)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("input end number: ")
	b, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	b = strings.TrimSpace(b)
	n2, _ := strconv.Atoi(b)

	if err != nil {
		log.Fatal(err)
	}

	for j := n1; j <= n2; j++ {
		if isPrime(j) {
			fmt.Printf("%d ", j)
		}
	}
}
