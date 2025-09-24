/* 함수 심화(Function Variable/First-class City)
¹ 콜백 함수(Callback Function): 다른 함수의 매개변수로 전달되어, 특정 조건이나 시점에 실행되는 함수
² First-class Citizen: 프로그래밍 언어에서 함수를 값처럼 취급할 수 있는 특성 (변수 할당, 매개변수 전달, 반환값 사용 가능)
³ 함수 타입(Function Type): Go에서 함수의 시그니처를 타입으로 정의하는 방법
⁴ 고차 함수(Higher-order Function): 함수를 매개변수로 받거나 함수를 반환하는 함수
*/
package main

import "fmt"

func multiply(x, y int) (r int) {
	r = x * y
	return
}

func sum(x, y int) (r int) {
	r = x + y
	return
}

func main() {
	//함수 고급
	// 1️⃣ 변수에 함수 할당 (Function Variable/First-class Function)

	//예제1 (슬라이스에 할당)
	f := []func(int, int) int{multiply, sum}

	a := f[0](10, 10)
	b := f[1](10, 10)

	fmt.Println("ex1 :", a, f[0](10, 10))
	fmt.Println("ex1 :", b, f[1](10, 10))
	fmt.Println()

	//예제2 (변수에 할당)
	var f1 func(int, int) int = multiply
	// func(int, int) int까지가 자료형을 표현한 것
	f2 := sum 

	fmt.Println("ex2 :", f1(10, 10))
	fmt.Println("ex2 :", f2(10, 10))
	fmt.Println()

	//예제3 (맵에 할당)
	m := map[string]func(int, int) int{
		"multiply": multiply,
		"sum":      sum,
	} // map[string]키는 문자열, value자리에 func가 온 것!
	fmt.Println("ex3 :", m["multiply"](10, 10))
	fmt.Println("ex3 :", m["sum"](10, 10))
}
