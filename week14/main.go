package main

import "fmt"

type student struct {
	id   int
	name string
	gpa  float32
}

func main() {
	var student1 student
	student1.id = 202444123
	student1.name = "이인희"
	student1.gpa = 4.42

	fmt.Println(student1.gpa)

	var student2 student
	student2.id = 202444456
	student2.name = "리인희"
	student2.gpa = 2.44

	fmt.Println(student2.id)

}
