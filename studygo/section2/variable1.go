// 변수1
package main

import "fmt"

func main() {
	/*
		디폴트 값으로 무조건 초기화 됨
		정수타입 : 0, 실수(소수점) : 0.0, 문자열 : "", Boolean : True, False
		변수명 : 숫자 첫글자x, 대소문자 구분o, 문자, 숫자, 밑줄, 특수기호 사용 가능
		변수 및 상수 : 함수 내외 사용하기 위해서 선언
	*/
	var a int
	var b string
	var c, d, e int
	var f, g, h int = 1, 2, 3
	var i float32 = 11.4
	var j string = "Hi! Golang!"
	var k = 4.75 // 선언 동시 초기화
	var l = "Hi! Seoul!"
	var m = true

	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
	fmt.Println("c : ", c)
	fmt.Println("d : ", d)
	fmt.Println("e : ", e)
	fmt.Println("f : ", f)
	fmt.Println("g : ", g)
	fmt.Println("h : ", h)
	fmt.Println("i : ", i)
	fmt.Println("j : ", j)
	fmt.Println("k : ", k)
	fmt.Println("l : ", l)
	fmt.Println("m : ", m)

}
