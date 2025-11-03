// 함수 심화(익명함수 : 즉시실행함수 실행 동시에 소멸함) [콜백함수나 클로져 구현에 많이 사용함]
package main

import "fmt"

func main() {
	//함수 고급
	//익명 함수(Anonymous Functions)
	//즉시 실행(선언과 동시에)

	//예제1
	func() {
		fmt.Println("ex1 : Anonymous func!")
	}() // 익명함수라서 다른데에서 호출이 불가, 그래서 맨 뒤에 호출하는 괄호를 넣음

	fmt.Println()

	//예제2
	msg := "Hello Golang!"

	func(m string) {
		fmt.Println("ex2 : ", m)
	}(msg)

	fmt.Println()

	//예제3
	func(x, y int) {
		fmt.Println("ex3 : ", x+y)
	}(10, 20) // 선언하지 말고 직접 값을 넣어주는 형태

	fmt.Println()

	//예제4
	r := func(x, y int) int {
		return x * y
	}(10, 10) // 선언하는 변수에 함수를 실행해서 리턴 값을 대입하는 형식
	fmt.Println("ex4 : ", r)
}
