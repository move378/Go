// 함수 기초(1)
package main

import (
	"fmt"
	"strconv"
)

// 함수 선언 위치는 어느 곳이든 가능
func helloGolang() {
	fmt.Println("ex1 : Hello Golang!")
}

/*
func helloGolang() //예외 발생(괄호 위치 컴파일 에러)
{
	fmt.Println("Hello Golang!")
}
*/

func say_one(m string) {
	fmt.Println("ex2 :", m)
}

func sum(x int, y int) int {
	return x + y
}

func main() {
	/*함수
	선언 : func 키워드로 선언 / 반환값의 유무와 매개변수의 유무
	func 함수명(매개변수) (반환타입 or 반환 값 변수명) : 반환 값 존재
	func 함수명() (반환타입 or 반환 값 변수명) : 매개변수 없음, 반환 값 존재
	func 함수명(매개변수) : 매개변수 존재, 반환 값 없음
	타 언어와 달리 반환 값(return value) 여러 개 가능 */

	//예제1
	helloGolang()

	//예제2
	say_one("Hello World!")

	//예제3
	result := sum(5, 5)
	fmt.Println("ex3 :", result)
	fmt.Println("ex3 :", sum(10, 10))
	fmt.Println("ex3 :", strconv.Itoa(sum(10, 10))) //int to string (strconv.Atoi(int) : string to int)

	
}
