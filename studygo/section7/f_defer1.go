// 함수 Defer(지연함수: 함수 생명주기의 마지막에 실행)
package main

import "fmt"

func ex_f1() {
	fmt.Println("f1 : start")
	defer ex_f2() //마지막에 호출
	fmt.Println("f1 : end")
}

func ex_f2() {
	fmt.Println("f2 : called")
}

func main() {
	//Defer 함수 실행(지연)
	//Defer를 호출한 함수가 종료되기 직전에 호출 된다. 예약해놨다가 나중에 호출됨. 끝날 때
	//타 언어의 Finally 문과 비슷
	//주로 리소스 반환 등에 사용, 열린 파일 닫기, Mutex 잠금 해제

	//예제1
	ex_f1()

	//예제2
	sayHello("Golang!")

	//예제3
	stack()
	// 자료구조 스택 : first in last out, last in first out
}

func sayHello(msg string) { // 익명 함수에서의 활용
	defer func() {
		fmt.Println(msg)
	}()

	func() {
		fmt.Println("Hi ")
	}()
}

func stack() {
	for i := 1; i <= 10; i++ {
		defer fmt.Println("ex1 : ", i)
	}
}
