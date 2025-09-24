// 사용자 정의 타입(2 형변환 - 엄격함)
package main

import "fmt"

type cnt int

func main() {
	//기본 자료형 사용자 정의 타입

	//예제1
	a := cnt(5)
	// 다른 언어에 있는 생성자 함수와 혼동[어떤 클래스명과 같은 이름의 함수를 생성자 함수]

	fmt.Println("ex1 : ", a)

	//예제2 : 자동으로 형 변환되지 않는다. 엄격함!
	var b cnt = 15

	fmt.Println("ex2 : ", b)
	//testConvertT(b) //예외 발생 (중요!) 사용자 정의 타입 <-> 기본 타입 : 매개변수 전달 시에 변환해야 사용 가능 (cnt(5), int(5))
	testConvertT(int(b)) // 부모가 같으니 형변환이 가능!

	testConvertD(b)
	testConvertD(cnt(10)) //사용 가능
	//testConvertD(int(b)) //예외 발생 (중요!) 사용자 정의 타입 <-> 기본 타입 : 매개변수 전달 시에 변환해야 사용 가능 (cnt(5), int(5))
}

func testConvertT(i int) {
	fmt.Println("ex 2 : (Default Type) : ", i)
}

func testConvertD(i cnt) {
	fmt.Println("ex 2 : (Custom Type) : ", i)
}
