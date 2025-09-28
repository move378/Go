// 인터페이스 고급(1 빈 인터페이스 복습)
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

	/* 인터페이스 활용(빈 인터페이스)
	빈 인터페이스를 함수로 전달받아서 타입을 신경쓰지 않고 코딩하는 것이 고 언어만의 독창적인 특징이다
	함수내에서 어떠한 타입이라도 유연하게 매개변수로 받을 수 있다. (모든 타입 지정 가능)

	빈 인터페이스는 함수 매개변수, 리턴 값, 구조체 필드 등에서 사용 가능 -> 어떤 타입으로도 변환이 가능하니까
	모든 타입을 나타내기 위해 빈 인터페이스를 사용
	Runtime에 결정되는 동적 타입으로 생각하면 쉽다. (Java의 Object 타입의 형변환)
	*/

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
