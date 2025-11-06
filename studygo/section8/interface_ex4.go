// 인터페이스 고급(4 switch를 이용하여 현재 데이터형을 반환받기, 타입 검사)
package main

import (
	"fmt"
)

func checkType(arg interface{}) {

	//데이터형을 알고 싶을 때, 빈인터페이스 명에 쩜찍고 arg.(type) 원래의 데이터형을 반환
	switch arg.(type) { // 마법의 스위치! Go Lang 에서는 Break 문이 없어도 그냥 빠져 나옴.
	case bool:
		fmt.Println("This is a bool : ", arg)
	case int, int8, int16, int32, int64:
		fmt.Println("This is an int : ", arg)
	case float64:
		fmt.Println("This is a float64 : ", arg)
	case string:
		fmt.Println("This is a string : ", arg)
	case nil:
		fmt.Println("This is nil : ", arg)
	default:
		fmt.Println("What is  this type? : ", arg)
	}
}

func main() {
	/* 실제 타입 검사 switch 사용
	빈 인터페이스는 어떠한 자료형도 전달 받을 수 있으므로, 타입체크를 통해 형 변환 후 사용 가능합니다.
	[런타임 에러 발생]
	*/

	//예제1
	checkType(true)
	checkType(1)
	checkType(22.542)
	checkType(nil)
	checkType("Hello Golang!")
}
