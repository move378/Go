// 구조체 기본(2 다양한 선언방법)
package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func (a Account) Calculate() float64 {
	return a.balance + (a.balance * a.interest)
}

func main() {
	/*
		🔑 핵심 규칙:

		1. 포인터 리시버 메소드가 있는 인터페이스
		   → 반드시 포인터 타입(*T)으로 할당
		   → &Account{...} 또는 new(Account) 사용

		2. 세 가지 선언 방법:
		   - new(Account):    포인터 반환, 초기화 불가
		   - &Account{...}:   포인터 반환, 초기화 가능
		   - Account{...}:    값 반환

		3. 메소드 호출 vs 인터페이스 할당:
		   - 메소드 호출: Go가 자동 변환 (편리)
		   - 인터페이스 할당: 메소드 세트 규칙 엄격 적용
	*/

	// 방법1: new + 필드별 할당
	var kim *Account = new(Account)
	kim.number = "245-901"
	kim.balance = 10000000
	kim.interest = 0.015

	// 방법2: & + 리터럴 초기화 (가장 권장)
	hong := &Account{
		number:   "245-902",
		balance:  15000000,
		interest: 0.04,
	}

	// 방법3: new (초기화 나중에)
	lee := new(Account)
	lee.number = "245-903"
	lee.balance = 13000000
	lee.interest = 0.025

	fmt.Println("ex1 : ", kim)
	fmt.Println("ex1 : ", hong)
	fmt.Println("ex1 : ", lee)
	fmt.Printf("ex2 : %#v\n", kim)
	fmt.Printf("ex2 : %#v\n", hong)
	fmt.Printf("ex2 : %#v\n", lee)

	fmt.Println()

	fmt.Println("ex3 : ", int(kim.Calculate()))
	fmt.Println("ex3 : ", int(hong.Calculate()))
	fmt.Println("ex3 : ", int(lee.Calculate()))
}
