// 함수 기초(2)
package main

import "fmt"

func sum(i int, f func(int, int)) { // 매개변수에 함수가 들어가는 경우 콜백함수
	f(i, 10) // add(1, 2)를 호출
}

func add(a, b int) { // A도 같은 인트형
	fmt.Println("ex1 :", a+b)
}

func multi_value(i int) {
	i = i * 10
}

func multi_reference(i *int) {
	*i *= 10 // *i = *i * 10
}

func main() {
	// 패턴: 1) 함수(콜백), 2) 참조 전달(call by reference), 3) 값 전달(call by value)

	//예제1 (콜백 호출)
	sum(10, add) //함수 전달

	//예제2 (call by value : 값만 전달)
	a := 100

	multi_value(a)
	fmt.Println("ex2 :", a) // a의 값에 변화없음
	// 매개 변수의 값은 복사되었다가 함수 사용이 끝나면 가비지 셀렉됨

	//예제3 (reference by value : 참조 전달)
	b := 100

	multi_reference(&b)
	// 참조 주소값을 전달, 받는 곳에서 포인터형으로 받고 역참조로 접근
	fmt.Println("ex3 :", b) // 원본 값이 변경

}
