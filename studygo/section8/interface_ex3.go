// 인터페이스 고급(3 Type Assertions : 연산을 수행하기 위해 타입검사,
// if문을 이용해서 원래의 타입으로 형변환을 하는 비즈니스 로직을 만들어야 함)
package main

import (
	"fmt"
	"reflect"
)

func main() {
	/* 타입 변환 (Type Assertions)
	실행(런타임) 시에는 인터페이스에 할당한 변수는 실제 타입으로 변환 후 사용해야 하는 경우
	인터페이스.(다시돌아갈타입) 형식 -> 형 변환(최초 타입으로만 가능)
	interfaceVal.(type) */

	//예제1
	var a interface{} = 15

	b := a
	c := a.(int)
	//d := a.(float64) //런타임 에러 발생 panic - 최초 타입으로만 형변환 가능

	fmt.Println("ex1 : ", a)
	fmt.Println("ex1 : ", reflect.TypeOf(a)) // 빈인터페이스를 받아서 최초의 타입을 확인하는 함수

	fmt.Println("ex1 : ", b)
	fmt.Println("ex1 : ", reflect.TypeOf(b))

	fmt.Println("ex1 : ", c)
	fmt.Println("ex1 : ", reflect.TypeOf(c))

	fmt.Println()

	//예제2(저장된 타입 실제 타입 검사)
	if v, ok := a.(int); ok { //해당 타입 값, 두 번째 값으로 타입 체크 Boolean 값이 반환됨.
		fmt.Println("ex2 : ", v, ok)
	}

}
