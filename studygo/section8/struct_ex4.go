// 구조체 심화(4 구조체 임베디드 패턴으로 구현하는 상속 및 오버라이드:상속이 아니라 is A 관계라고 함[~도 A이다])
package main

import "fmt"

type Employee struct {
	name   string
	salary float64
	bonus  float64
}

func (e Employee) Calculate() float64 { // 리시버가 Employee로 잡혀있어도 Executives에서 사용가능[메소드 재사용]
	return e.salary + e.bonus
}

type Executives struct { // 구조체 안에 구조체를 넣어서 Employee > Executives 집합처럼 포함하는 안에 있는 구조
	Employee     // is a 관계 - Executives is employee 임원도 직원이다 [상속개념]
	specialBonus float64
}

func main() {
	//구조체 임베디드 패턴
	//다른 관점으로 메소드를 재 사용하는 장점 제공

	//상속을 허용하지 않는 Go 언어에서 메소드 상속 활용을 위한 패턴

	//예제1
	//직원
	ep1 := Employee{"kim", 2000000, 300000}
	ep2 := Employee{"park", 1500000, 200000}
	//임원
	ex := Executives{
		Employee{"lee", 4000000, 1000000},
		1000000,
	}

	fmt.Println("ex1 : ", int(ep1.Calculate()))
	fmt.Println("ex1 : ", int(ep2.Calculate()))
	//Employee 상속받은 부모를 통해서 메소드 호출(Go에서는 상속이 없기에 임베디드 '~에 속한'라고 하고 메서드 재사용이라 함)
	//(e Employee) 리시버를 통해서 아래처럼 바로 사용이 가능, 임원도 직원이니까!
	fmt.Println("ex1 : ", int(ex.Calculate()+ex.specialBonus))

}
