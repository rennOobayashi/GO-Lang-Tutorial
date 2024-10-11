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
	fmt.Print("Input ur score: ")
	in := bufio.NewReader(os.Stdin)
	i, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i)
	score, _ := strconv.ParseInt(i, 16, 32)
	fmt.Printf("%d ", score)

	if score >= 60 {
		fmt.Println("a")
	} else {
		fmt.Println("BCDF")
	}

}
