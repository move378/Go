// 구조체 기본(5 필드 태그(레이블) 사용하기! 및 상세 확인)
package main

import (
	"fmt"
	"reflect"
)

type Car struct { // 대문자로 생성해야 외부에서 참조가 가능하다 public
	name    string "차량명"
	color   string "색상"
	company string "제조사"
}

func main() {
	//필드 태그 사용

	//예제1
	tag := reflect.TypeOf(Car{}) // 태그를 지정한 뒤, 리플랙트 패키지의 TypeOf(빈 인터페이스)로 가져옴

	for i := 0; i < tag.NumField(); i++ { // Num.Field() 필드가 몇개인지를 반환하는 메소드
		fmt.Println("ex1 : ", tag.Field(i).Tag, tag.Field(i).Name, tag.Field(i).Type) // .Field(i)로 가져옴
	}
}
