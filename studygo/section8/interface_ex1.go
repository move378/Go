//인터페이스 고급(1)
package main

import "fmt"

type Dog struct {
	name   string
	weight int
}

type Cat struct {
	name   string
	weight int
}

func printValue(s interface{}) {
	fmt.Println(s)
}

func main() {
	//인터페이스 활용(빈 인터페이스)
	//함수내에서 어떠한 타입이라도 유연하게 매개변수로 받을 수 있다. (모든 타입 지정 가능)
	//빈 인터페이스 : 함수 매개변수, 리턴 값, 구조체 필드 등으로 사용

	//예제1
	dog := Dog{"poll", 10}
	cat := Cat{"bob", 5}

	//빈 인터페이스 : 어떤 값이든 허용 가능(유연성 증가)
	printValue(dog)
	printValue(cat)
	printValue(15)
	printValue("Animal")
	printValue(25.5)
	printValue([]Dog{})

}
