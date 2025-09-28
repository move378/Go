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
		&struct, struct : &struct 포인터를 받아오기, 역참조를 또 하기 때문에 속도는 조금 느리다.
		포인터형으로 구조체를 사용해야 하는 경우 - 인터페이스 메소드를 선언만 해둔 후
		-> 오버라이딩(재정의)한 메서드에 포인터 리시버를 사용할 경우 반드시 &struct 형식으로 넘겨야 작동한다.

		.\main.go:37: cannot use struct (type structType) as tyep MyInterface in argument to DisplayInfo:
		structType does not implement MyInterface (DisplayA method has pointer receiver)
	*/
	//예제1
	//선언 방법1
	var kim *Account = new(Account)
	// 인스턴스를 생성해서 변수에 참조를 전달해야 하기 때문에 주소값 표시 *를 쓴다.
	// 인터페이스를 오버라이딩 할 때 new라는 키워드를 사용해야 한다.
	kim.number = "245-901"
	kim.balance = 10000000
	kim.interest = 0.015

	//선언 방법2
	hong := &Account{number: "245-902", balance: 15000000, interest: 0.04}

	//선언 방법3
	lee := new(Account) // new 키워드가 쓰이면 선언만 가능!
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

	//예제2
	fmt.Println("ex3 : ", int(kim.Calculate()))
	fmt.Println("ex3 : ", int(hong.Calculate()))
	fmt.Println("ex3 : ", int(lee.Calculate()))

}
