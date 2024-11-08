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

func getInteger() (int, error) {
	in := bufio.NewReader(os.Stdin)
	i, err := in.ReadString('\n')

	if err != nil {
		return 0, err
	}

	i = strings.TrimSpace(i)
	n, _ := strconv.Atoi(i)

	if err != nil {
		return 0, err
	}

	return n, nil
}

func main() {

	fmt.Print("input start number: ")
	n1, err1 := getInteger()
	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Print("input end number: ")
	n2, err2 := getInteger()
	if err2 != nil {
		log.Fatal(err2)
	}

	for i := n1; i <= n2; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
}
