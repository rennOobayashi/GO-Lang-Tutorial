package greeting

import "fmt"

// 대문자로 시작하지 않으면 외부에서 쓸수 없음
func hi(name string) {
	fmt.Printf("Hi %s!\n", name)
}

func hello(name string) {
	fmt.Printf("Hello %s!\n", name)
}

func EnglishGreetings(name string) {
	hi(name)
	hello(name)
}
