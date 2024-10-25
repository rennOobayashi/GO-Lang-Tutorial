package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var win bool = false

	rand.Seed(time.Now().Unix())
	answer := rand.Intn(6)
	fmt.Println(answer)

	for guesses := 0; guesses < 3; guesses++ {
		fmt.Printf("숫자입력(남은 기회 %d회): ", 3-guesses)
		in := bufio.NewReader(os.Stdin)
		i, err := in.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		i = strings.TrimSpace(i)
		guess, _ := strconv.Atoi(i)

		if answer == guess {
			fmt.Println("입력하신 수는 정답입니다.")

			win = true
			break
		} else if answer > guess {
			fmt.Println("입력하신 수는 정답보다 작은 수 입니다.")
		} else {
			fmt.Println("입력하신 수는 정답보다 큰 수 입니다.")
		}
	}

	if win {
		fmt.Println("\n********************\n*당신이 이겼습니다.*\n********************")
	} else {
		fmt.Println("\n------------------\n|당신이 졌습니다.|\n------------------")
	}

}
