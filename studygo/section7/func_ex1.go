// 함수 심화(가변인자)
package main

import "fmt"

func multiply(n ...int) int {
	tot := 1
	for _, value := range n {
		tot *= value
	}

	return tot
}

func sum(n ...int) int {
	tot := 0
	for _, value := range n {
		tot += value
	}

	return tot
}

func prtWord(msg ...string) {
	for _, value := range msg {
		fmt.Println("ex2 :", value)
	}
}

func main() {
	//함수 고급
	//가변 인자 실습(매개변수 개수가 동적으로 변할 때 - 정해져있지 않음)
	//자바스크립트 ES6 문법에서도 쓰이는 문법

	//예제1
	x := multiply(5, 6, 7, 8, 9, 10)
	y := sum(5, 6, 7, 8, 9, 10)
	fmt.Println("ex1 :", x)
	fmt.Println("ex1 :", y)
	fmt.Println()

	//예제2
	prtWord("a", "apple", "test", "seoul", "golang", "hi")
	fmt.Println()

	//예제3
	a := []int{5, 6, 7, 8, 9, 10} // 크기가 정해져있지 않으니 슬라이스 형

	m := multiply(a...) // 스프레드 연산자
	n := sum(a...)

	fmt.Println("ex3 :", m)
	fmt.Println("ex3 :", n)
}
