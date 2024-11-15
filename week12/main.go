package main

import (
	"fmt"
	"time"
)

func main() {

	var scores [3]int

	scores[1] = 90
	fmt.Println(scores[1], scores[0], scores[2])
	fmt.Printf("%#v\n", scores)
	fmt.Println(scores)

	// var dates [3]time.Time

	// dates[1] = time.Unix(198828330, 0)

	// fmt.Println(dates[1])
	// var scores [3]int

	// scores[1] = 90

	// fmt.Println(scores[1], scores[0], scores[2])

	dates := [3]time.Time{
		time.Unix(0, 0),
		time.Unix(1, 0),
		time.Unix(194891203, 0), //need comma
	}

	fmt.Println(dates[0], dates[1], dates[2])
	fmt.Println(dates)

	// fmt.Printf("%#v\n", dates)
}
