// 인터페이스 기본(2 인터페이스의 구현)
package main

import "fmt"

type Dog struct {
	name   string
	weight int
}

// bite 메소드 구현
func (d Dog) bite() {
	fmt.Println(d.name, " bites!")
}

// 동물의 행동 인터페이스 선언
type Behavior interface {
	bite()
}

func main() {

	// 예제1: 인터페이스 변수 선언 후 할당
	dog1 := Dog{"poll", 10}
	var inter1 Behavior // inter1은 nil (zero value)
	inter1 = dog1       // 암시적 타입 변환: Dog → Behavior
	inter1.bite()       // 다형성: 런타임에 Dog.bite() 호출

	// 예제2: 인터페이스 선언과 동시에 초기화
	dog2 := Dog{"mary", 12}  // Dog 타입 인스턴스 생성
	inter2 := Behavior(dog2) // dog2를 Behavior 인터페이스 타입으로 변환(타입 어서션)
	//dog2의 메모리 주소와 타입 정보가 inter2 인터페이스에 저장됨
	inter2.bite()

	// 예제3: 인터페이스 슬라이스 (컬렉션)
	inters := []Behavior{dog1, dog2}

	// 방법1: 인덱스 기반 순회
	for idx, _ := range inters {
		//dog1.bite() Behavior타입으로 받아서 배열에 넣고 실행하는 것임
		inters[idx].bite()
	}

	// 방법2: 값 기반 순회 (더 관용적)
	// 해당 값은 구조체 객체 자체를 반환할 테니까,
	for _, val := range inters {
		// inter2 := Behavior(dog2) 와 같음
		val.bite()
	}

	// 추가 학습: 인터페이스 내부 구조 이해
	fmt.Printf("inter1 타입: %T, 값: %v\n", inter1, inter1)
	fmt.Printf("inter2 타입: %T, 값: %v\n", inter2, inter2)

}
