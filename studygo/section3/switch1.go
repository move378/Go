// Switch문 (1)

package main

import "fmt"

func main() {
	/* 제어문(조건문) - switch
	switch 뒤 표현식(Expression) 생략 가능
	case 뒤 표현식(Expression) 	사용 가능
	자동 break 때문에 fallthrouth 존재
	Type 분기 -> 값이 아닌 변수 Type로 분기 가능

	어떤 범위로 맵핑을 하는거면 if, else if, else 로 쓰고
	값이 딱딱 떨어지는 듯이 정해져 있을 떄, switch case를 쓰면 스위치 처럼 사용할 수 있으니까 더 유용하다.
	스위치가 사실 if 문보다 다양하게 사용이 가능하다
	*/

	// 예제 1
	a := -7
	switch {
	case a < 0:
		fmt.Println(a, "는 음수")
	case a == 0:
		fmt.Println(a, "는 0")
	case a > 0:
		fmt.Println(a, "는 양수")
	}

	//예제2 어떠한 영역 안에서 쓸 때는 짧은 선언을 선호 합니다. 범위 Scope가 제한이 걸리는 것으로 Clean Code !!
	switch b := 27; {
	case b < 0:
		fmt.Println(b, "는 음수")
	case b == 0:
		fmt.Println(b, "는 0")
	case b > 0:
		fmt.Println(b, "는 양수")
	}

	// 예제3 패턴1
	switch c := "go"; c { // 여기서는 표현식이 없는 경우 같은 것을 매칭 시킴
	case "go":
		fmt.Println("Go!")
	case "java":
		fmt.Println("Go!")
	default:
		fmt.Println("일치하는 값 없음")
	}

	// 예제4 패턴2 같은 타잎
	switch c := "go"; c + "lang" {
	case "golang":
		fmt.Println("Golang!")
	case "java":
		fmt.Println("Java")
	default:
		fmt.Println("일치하는 값 없음")
	}

	//예제5 패턴3 짧은 선언
	switch i, j := 20, 30; {
	case i < j:
		fmt.Println("i는 j보다 작다")
	case i == j:
		fmt.Println("i는 j보다 같다")
	case i > j:
		fmt.Println("i는 j보다 크다")
	}

}
