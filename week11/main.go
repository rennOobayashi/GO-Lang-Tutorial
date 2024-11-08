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

func getInteger() int {
	in := bufio.NewReader(os.Stdin)
	i, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i)
	n, _ := strconv.Atoi(i)

	if err != nil {
		log.Fatal(err)
	}

	return n
}

func main() {
	var n1, n2 int

	fmt.Print("input start number: ")
	n1 = getInteger()

	fmt.Print("input end number: ")
	n2 = getInteger()

	for i := n1; i <= n2; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
}
