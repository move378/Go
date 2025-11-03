// 데이터 타입 : 문자열 연산(1)
package main

import (
	"fmt"
	"strings"
)

func main() {
	/* 문자열 연산
	문자열에서 특정문자 열을 [추출],
	문자열이 같은지 다른지를 [비교],
	a문자열과 b문자열을 [조합]하거나 [결합].
	문자열을 다룬다는 것은 성능에 중요한 이슈! 이부분에 많은 개발자들이 고민을 합니다.
	*/

	//예제1(추출)
	var str1 string = "Golang"
	var str2 string = "World"

	fmt.Println("ex1 : ", str1[0:2], str1[0]) // 0~2번 슬라이싱 문자로 개별 인덱스로 가져오면 정수형 출력!
	// 이때 특이하게도 슬라이스가 인덱스 0~2-1 까지 출력됨 Go 까지만 출력됨
	fmt.Println("ex1 : ", str2[3:], str2[0]) //[3]인덱스부터 ~비워놓으면 인덱스의 끝까지
	fmt.Println("ex1 : ", str2[:4], str2[0]) //[4]인덱스까지인데 항상 끝부분은 -1로 계산 [서양식]
	fmt.Println("ex1 : ", str1[1:3])

	//예제2(비교)
	str3 := "Golang"
	str4 := "World"

	fmt.Println("ex2 : ", str3 == str4) //바이트로 비교
	fmt.Println("ex2 : ", str3 != str4)
	fmt.Println("ex2 : ", str3 > str4)
	fmt.Println("ex2 : ", str3 < str4)
	// Go 문자열 -> 아스키 코드에 대한 사전식 비교

	/*
		문자열 조합은 한 번 생성 후 수정 불가 이유로 새로 계속해서 생성 된다.
		효율적 사용을 위해서 되도록 join 함수 사용 추천 [성능 이슈 : 중요]
	*/

	//예제3(결합 : string append)
	str5 := "Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to " +
		"write programs that get the most out of multicore and networked machines, while its novel " +
		"type system enables flexible and modular program construction. Go compiles quickly to machine " +
		"code yet has the convenience of garbage collection and the power of run-time reflection."
	str6 := "It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language."

	fmt.Println("ex3 : ", str5+str6)

	//예제2(결합 : Join을 써서 내부 연산을 최소화 하는 상황)
	strSet := []string{}          //슬라이스 선언
	strSet = append(strSet, str5) // append는 import "string" 에 있음
	strSet = append(strSet, str6)
	fmt.Println("ex4 : ", strings.Join(strSet, "--구분 매개변수--"))

	// 결론 조인을 사용하거나, 성능 좋은 서드 파티 패키지를 이용해서 string 연산을 할 것!
}
