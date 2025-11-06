// 인터페이스 기본(4 익명 인터페이스:타입 정의 생략으로 타입이 익명타입!)
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

func (d Dog) run() {
	fmt.Println(d.name, "Dog is running!")
}

func (c Cat) run() {
	fmt.Println(c.name, "Cat is running!")
}

// 익명 인터페이스(타입 정의x)
func act(animal interface{ run() }) {
	// 익명 인터페이스 : 따로 선언하지 않고 파라미터에서 즉시 구현해도 run() 메소드를 구현하였으면 덕타이핑 적용!
	// Behavior 타입명을 사용하지 않고 변수명(객체명)으로 즉시 인터페이스 구현. 구지 따로 사용자 정의 타입을 선언하지 않음.
	animal.run() // run()의 구현여부만 판별해서 타입을 체크
}

func main() {
	//익명 인터페이스 사용 예제(즉시 선언 후 사용)

	//예제1
	dog := Dog{"poll", 10}
	cat := Cat{"bob", 5}

	//개 행동 실행
	act(dog)
	//고양이 행동 실행
	act(cat)

}
