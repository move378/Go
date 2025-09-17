//IF문(1)

package main

import "fmt"

func main() {
	/* 제어문(조건문)
	   IF 문 : 반드시 Boolean 검사 -> 1, 0으로 안됨. 명시적으로 true인지 false인지 정확히 해야 함(자동 형 변환 불가)
	   소괄호 사용 x
	*/

	var a int = 20
	b := 20

	// 예제1
	if a >= 15 {
		fmt.Println("15 이상")
	}

	if b >= 25 {
		fmt.Println("25 이상")
	}
	/* 에러상황
	   1) a >= 15; 이게 선언하고 들어가는 것인데, 고에서는 엔터가 컴파일할 때 ; 이다. 그래서 괄호 위치 유의
	   바로 { } 열어줘야지 엔터를 찍고 열면 a >=15; {} 이런식으로 코딩한 것.
	   2) 고에서는 한줄이라도 {} 대괄호를 써줘야 합니다.
	   3) if c:= 1; boolean 값이여야 하고, c { fmt.Println("True") }

	*/

	if c := true; c {
		fmt.Println("True")
	}

	if c := 40; c >= 35 {
		fmt.Println("35 이상")
	}
	// c += 20 짧은 선언은 {} 안에서만 일회용으로 사용되고 없어진다.
	// 그리고 if문 바깥에서도 main 함수 안이기 때문에 짧은 선언을 사용할 수 있다.
}
