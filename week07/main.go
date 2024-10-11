package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	year := now.Year()
	month := int(now.Month())
	days := now.Day()

	//오늘은 2024년 10월 11일 입니다.
	fmt.Printf("오늘은 %d년 %d월 %d일 입니다.\n", year, month, days)
	fmt.Printf("지금은 %d시 %d분 %d초 입니다.\n", now.Hour(), now.Minute(), now.Second())
}
