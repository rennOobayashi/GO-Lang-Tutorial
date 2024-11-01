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

	cnt := 0
	j := 2
	for j < n {
		if n%j == 0 {
			cnt++
		}

		j++
	}

	if cnt == 0 {
		fmt.Printf("%d is prime number.\n", n)
	} else {
		fmt.Printf("%d is not prime number.\n", n)
	}
}
