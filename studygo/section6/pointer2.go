// 자료형 : 포인터(3)
package main

import "fmt"

func rptc(n *int) { // 파라미터에서 int의 포인터를 전달 받겠다고 함
	*n = 77 // 파라미터로 받은 포인터의 실제 값에 77을 대입
}

func vptc(n int) int {
	n = 77 // 복사본만 77로 변경
	return n
}

func main() {
	/* 포인터 값 전달
	함수, 메서드 호출 시에 매개변수 값을 얕은 복사를 통해서 전달(참조만) -> 함수, 메서드 내에서는 원본 값 변경 불가능
	해당하는 원본 값의 변경 위해서 포인터를 사용해 변경 가능
	특히 크기가 큰 배열[값]인 경우 값 복사시에 시스템 부담 -> 포인터 전달로 참조형식으로 해결 But (슬라이스는 참조 전달) */

	//예제1
	var a int = 10
	var b int = 10
	fmt.Println("ex1 : ", a)
	fmt.Println("ex1 : ", b)
	fmt.Println()

	rptc(&a)          // a는 주소값을 전달해서 원본 a가 변경됨
	result := vptc(b) // b는 주소값을 전달한 것이 아니고, 복사해서 넘김 원본에 변경없음
	/*
		매개변수로 b를 value 값 10을 복사해서 넘김, b의 복사본을 매개변수로 넘기고,
		해당 함수에서 b의 복사본 n에 77을 대입, 실제 b의 값은 변동 없음

		Pass by Value: 값에 의한 전달 (복사)
			데이터 값을 복사해서 전달
			복사본에만 접근
			원본 데이터 안전하게 보호
			작은 데이터에 적합
		Pass by Pointer: 포인터에 의한 전달 (주소)
			메모리 주소를 전달
			원본 데이터에 직접 접근
			함수 내에서 원본 변경 가능
			메모리 효율적 (큰 데이터 구조체)
		Dereference: 역참조 (*p로 포인터가 가리키는 값에 접근)
		Address-of: 주소 연산자 (&변수로 변수의 주소 획득)
	*/

	//vptc(&b) //에러 발생

	fmt.Println("ex1 : ", a)
	fmt.Println("ex1 : ", b)
	fmt.Println("ex1 : ", result)
}
