// 사용자 정의 타입(1 메소드 바인딩)
package main

import "fmt"

// 사용자 정의 타입(구조체)
type Car struct {
	name  string
	color string
	price int64
	tax   int64
}

// 구조체 <-> 메소드 바인딩 [어떤 구조체(속성과 상태)에 기능을 추가하는 것]
// 구조체에 대한 주소 값이 있고, 그와 연결된 메소드가 있는 형태인가?
func (c Car) Price() int64 {
	return c.price + c.tax
}

func main() {
	/*Go -> 객체 지향 타입을 구조체로 정의한다.(클래스, 상속 개념 없음)
	객체 지향 -> 클래스 (속성 상태값 : 명사)) : 멤버변수, 기능(메소드:동사) - 코드의 재사용성, 관리 용이, 신뢰성
	Go는 전형적인 객체지향의 특징을 가지고 있지 않지만, 사용자 정의 타입 = 구조체
	인터페이스 -> 다형성 지원, 구조체로써 클래스 형태의 코딩이 가능
	객제지향의 기본 개념 -> Go에서 포함하고 있다. -> 객체 지향 프로그래밍 언어

	Go Lang 의 특징
	상태, 메소드를 분리해서 정의(결합성 없음)
	사용자 정의 타입 : 구조체 , 인터페이스 , 기본타입(int, float, string...) , 함수 등 오버라이딩 가능
	구조체와 -> 메서드 연결을 통해서 타 언어의 클래스 형식 처럼 사용 가능(객체 지향)
	*/
	//예제1
	bmw := Car{name: "520d", price: 54500000, color: "white", tax: 545000}
	benz := Car{name: "220d", price: 74500000, color: "white", tax: 745000}

	bmwPrice := bmw.Price()
	benzPrice := benz.Price()

	fmt.Println("ex1 : ", &bmw)
	fmt.Println("ex1 : ", bmwPrice)
	fmt.Println("ex1 : ", bmw.Price())

	fmt.Println("ex1 : ", &benz == &bmw)

	fmt.Println("ex1 : ", &benz)
	fmt.Println("ex1 : ", benzPrice)
	fmt.Println("ex1 : ", benz.Price())

}
