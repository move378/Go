//구조체 기본(6)
package main

import "fmt"

type spec struct { //소문자로 선언
	length int "전장"
	height int "전고"
	width  int "전축"
}

type Car struct { //대문자로 선언
	name    string
	color   string
	company string
	detail  spec
}

func main() {
	//중첩 구조체

	//예제1
	car1 := Car{
		"520d",
		"silver",
		"bmw",
		spec{4000, 1000, 2000},
	}
	// 내부 spec 구조체 값 출력
	fmt.Println("ex1 : ", car1.name)
	fmt.Println("ex1 : ", car1.color)
	fmt.Println("ex1 : ", car1.company)
	fmt.Printf("ex1 : %#v\n", car1.detail)
	fmt.Println()

	//예제2
	// 내부 spec 구조체 필드 값 출력
	fmt.Println("ex2 : ", car1.detail.length)
	fmt.Println("ex2 : ", car1.detail.height)
	fmt.Println("ex2 : ", car1.detail.width)

}
