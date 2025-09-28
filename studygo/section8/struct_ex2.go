// 구조체 심화(2 함수의 매개변수에 포인터가 오는 경우)
package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func CalculateD(a Account) { //값 복사 전달
	a.balance = a.balance + (a.balance * a.interest)
}

func CalculateP(a *Account) { //주소(참조) 전달 [포인터형임으로 &로 주소를 전달해야 함]
	a.balance = a.balance + (a.balance * a.interest)
}

func main() {
	//구조체 생성자 패턴 예제

	//예제1
	kim := Account{number: "245-901", balance: 10000000, interest: 0.015}
	lee := Account{number: "245-902", balance: 12000000, interest: 0.045}

	fmt.Println("ex1 : ", kim)
	fmt.Println("ex1 : ", lee)
	fmt.Println()

	CalculateD(kim)
	CalculateP(&lee) //CalculateP(lee) 호출 시 예외 발생

	fmt.Println("ex1 : ", int(kim.balance))
	fmt.Println("ex1 : ", int(lee.balance))

}
