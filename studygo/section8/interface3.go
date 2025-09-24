//인터페이스 기본(3)
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

//구조체 Dog 메소드 구현
func (d Dog) bite() {
	fmt.Println("Dog bites!")
}

func (d Dog) sounds() {
	fmt.Println("Dog barks!")
}

func (d Dog) run() {
	fmt.Println("Dog is running!")
}

//구조체 Cat 메소드 구현
func (c Cat) bite() {
	fmt.Println("Cat bites!")
}

func (c Cat) sounds() {
	fmt.Println("Cat cries!")
}

func (c Cat) run() {
	fmt.Println("Cat is running!")
}

//인터페이스를 파라미터로 받는다.
func act(animal Behavior) {
	animal.bite()
	animal.sounds()
	animal.run()
}

//동물의 행동 인터페이스 선언
type Behavior interface {
	bite()
	sounds()
	run()
}

func main() {
	//인터페이스 구현 예제
	//인터페이스 규격화 역할 이해
	//인터페이스에 정의된 메소드 사용 유도
	//코드의 가독성 및 유지보수 증가

	//덕타이핑 예제
	//덕타이핑 : 구조체 및 변수의 값이나 타입은 상관하지 않고 오로지 구현된 메소드로만 판단하는 방식
	//Go의 중요한 특징 : 오리처럼 걷고, 소리내고, 헤엄 등 행동이 같으면 오리라고 볼 수 있다 -> 의미

	//예제1
	dog := Dog{"poll", 10}
	cat := Cat{"bob", 5}

	//개 행동 실행
	act(dog)
	//고양이 행동 실행
	act(cat)

}
