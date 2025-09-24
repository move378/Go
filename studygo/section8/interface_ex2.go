//인터페이스 고급(2)
package main

import "fmt"

func main() {
	//빈 인터페이스 타입 상세 설명
	//모든 타입을 나타내기 위해 빈 인터페이스 사용
	//동적타입으로 생각하면 쉽다(타 언어의 Object 타입)

	//예제1
	var a interface{}

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
	fmt.Printf("Type : (%T) ", v) //원본 타입
	fmt.Println("ex1 : ", v)
}
