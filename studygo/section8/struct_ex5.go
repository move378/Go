// 구조체 심화(5 구조체 임베디드 - 메소드 오버라이딩[제정의])
package main

import "fmt"

type Employee struct {
	name   string
	salary float64
	bonus  float64
}

func (e Employee) Calculate() float64 {
	return e.salary + e.bonus
}

// 이름이 같은 함수 사용 : 오버라이딩
func (e Executives) Calculate() float64 {
	return e.Employee.salary + e.Employee.bonus + e.specialBonus
}

type Executives struct {
	Employee
	specialBonus float64
}

func main() {
	//구조체 임베디드 메소드 오버라이딩 패턴
	//
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

	fmt.Println("ex1 : ", int(ex.Calculate()+ex.specialBonus))          //오버라이딩 : 잘못 된 값 반환
	fmt.Println("ex1 : ", int(ex.Calculate()))                          //오버라이딩 : 정확한 값
	fmt.Println("ex1 : ", int(ex.Employee.Calculate()+ex.specialBonus)) // 임베디드된 원래의 Calculate 호출

}
