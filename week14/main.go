package main

import "fmt"

func main() {
	ages := make(map[string]int)

	var name string
	var age int

	for {
		fmt.Print("Enter your name(quit to q): ")
		fmt.Scanln(&name)
		if name == "q" || name == "Q" {
			break
		}

		fmt.Print("Enter your age: ")
		fmt.Scanln(&age)

		ages[name] = age
	}

	for name, age := range ages {
		fmt.Printf("%s is %d years old.\n", name, age)
	}
}
