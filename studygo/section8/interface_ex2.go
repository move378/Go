// 인터페이스 고급(2 빈인터페이스를 선언해서 변수처럼 사용하는 것(빈 인터페이스 형태의 펑션 함수형 프로그래밍임.))
package main

import "fmt"

func main() {
	//빈 인터페이스 타입 상세 설명
	//모든 타입을 나타내기 위해 빈 인터페이스 사용
	//동적타입으로 생각하면 쉽다(타 언어의 Object 타입)

	//예제1
	var a interface{}
	// 요게 자료형을 명시하지 않은 빈인터페이스를 변수로써 선언, 함수형 프로그래밍
	// 이렇게 하면 변수에 모든 타입을 넣을 수 있음.
	// 타입 어써션, 다시 원본 타입으로 형변환을 할 수 있음

	printContents(a)

	a = 7.5
	printContents(a)

	a = "Golang"
	printContents(a)

	a = true
	printContents(a)

	a = nil
	printContents(a)
}

func printContents(v interface{}) {
	fmt.Printf("Type : (%T) ", v) //원본 타입 : 이게 나중에 연산을 하려면 형변환을 해서 연산을 해야 함.
	fmt.Println("ex1 : ", v)
}
