package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/keyboard"
)

func main() {
	var gpa [3]float64

	for i := 0; i < len(gpa); i++ {
		fmt.Printf("input float number: ")
		s, err := keyboard.GetFloat()

		if err != nil {
			log.Fatal(err)
		}

		gpa[i] = s

	}

	for _, score := range gpa {
		fmt.Printf("%f ", score)
	}
}
