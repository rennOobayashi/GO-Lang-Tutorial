package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("input number: ")
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

	isPrime := true

	//1 is not prime number
	if n <= 1 {
		isPrime = false
	} else {
		j := 2
		for j < n {
			if n%j == 0 {
				isPrime = false

				break //performance up
			}

			j++
		}
	}

	fmt.Printf("\n")

	if isPrime {
		fmt.Printf("%d is prime number.\n", n)
	} else {
		fmt.Printf("%d is not prime number.\n", n)
	}
}
