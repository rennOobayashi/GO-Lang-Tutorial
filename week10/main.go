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

	//cnt := 0
	isPrime := true //가독성과 메모리 효율성 증가
	j := 2
	for j < n {
		if n%j == 0 {
			//cnt++
			isPrime = false //더하기 연사 제거
		}

		j++
	}

	if isPrime { //== 비교 연산 제거
		fmt.Printf("%d is prime number.\n", n)
	} else {
		fmt.Printf("%d is not prime number.\n", n)
	}
}
