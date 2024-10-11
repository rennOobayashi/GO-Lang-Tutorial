package main

import (
	"fmt"
	"strings"
)

func main() {
	var army string = "우리는 $가와 $민에 충성을 다하는 대한민$ 육군이다."
	armyfixed := strings.NewReplacer("$", "국")
	fmt.Println(army)
	fmt.Println(armyfixed.Replace(army))

	// var now time.Time = time.Now()
	// year := now.Year()
	// month := int(now.Month())
	// days := now.Day()

	// //오늘은 2024년 10월 11일 입니다.
	// fmt.Printf("오늘은 %d년 %d월 %d일 입니다.\n", year, month, days)
	// fmt.Printf("지금은 %d시 %d분 %d초 입니다.\n", now.Hour(), now.Minute(), now.Second())
}
