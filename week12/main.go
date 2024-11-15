package main

import (
	"fmt"
	"time"
)

func main() {
	var scores [3]int

	scores[1] = 90

	//for i := 0; i < len(scores); i++ //<-panic: runtime error: index out of range [3] with length 3
	for i := 0; i < len(scores); i++ { //unsafe
		fmt.Printf("%d ", scores[i])
	}

	fmt.Println()

	dates := [3]time.Time{
		time.Unix(0, 0),
		time.Unix(1, 0),
		time.Unix(194891203, 0), //need comma
	}

	// fmt.Println(dates[0], dates[1], dates[2])
	// fmt.Println(dates)

	// for i, date := range dates {
	// 	fmt.Println(i, date)
	// }

	for _, date := range dates { //like python, safe
		fmt.Println(date)
	}
}
