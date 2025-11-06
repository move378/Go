// 구조체 심화(3 포인터 리시버를 사용하는 경우) - 포인터 리시버로 구현하는 경우 & 기호 사용안해도 됨
package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

//리시버 변수 사용안할 경우 func(_ Account) -> 밑줄로 생략 가능

func (a Account) CalculateD(bonus float64) { //값 복사 전달
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func (a *Account) CalculateP(bonus float64) { //주소(참조) 전달
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func main() {
	//구조체 생성자 패턴 예제

	//정리 : 구조체 인스턴스 값 변경 시 -> 포인터(참조) 전달 , 보통의 경우 -> 값(복사) 전달

	//예제1
	kim := Account{number: "245-901", balance: 10000000, interest: 0.015}
	lee := Account{number: "245-902", balance: 10000000, interest: 0.045}

	fmt.Println("ex1 : ", kim)
	fmt.Println("ex1 : ", lee)
	fmt.Println()

	lee.CalculateD(1000000)
	kim.CalculateP(1000000)

	fmt.Println("ex1 : ", int(lee.balance))
	fmt.Println("ex1 : ", int(kim.balance))

}
