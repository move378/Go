// 함수 기초(2)
package main

import "fmt"

func sum(i int, f func(int, int)) {
	/* 매개변수에 함수가 들어가는 경우 콜백함수(lambda 로 많이 들어감) 콜백의 개념이 어려운 것이 아니라,
	[매개변수로 함수를 받는다는 의미를 생각해보면 함수를 받아서 그 함수를 명령구문(코드블럭)에서 함수를 실행한다는 뜻],
	함수가 중첩되어 있는 구조로 뒤에 있는 함수를 먼저 실행해서 결과를 얻겠다는 의미임 */
	f(i, 10)
	/* add(i, 10)를 호출하니까 이건 더하는 기능을 함수의 매개변수로 받는 함수임.
	정리하면 어떠한 함수로 정의된 메소드나 기능을 전달받을 수 있다는 것 */
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
	// 패턴: 1) 함수(콜백), 2) 값 전달(call by value), 3) 참조 전달(call by reference)

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
