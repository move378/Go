// 사용자 정의 타입(4)
package main

import "fmt"

// 타입정의 장바구니 구조체{ 매개변수 형식 }
type shopingBasket struct{ cnt, price int }

//구매 함수
func (b shopingBasket) purchase() int { // 이거 자체가 b의 구조체를 받는 purchase() 함수니까 메소드!
	return b.cnt * b.price
}

//원본 수정 (참조 형식) : 형식만 지정해주면 * 넣어서 역참조하고 그럴 필요 없음
func (b *shopingBasket) rePurchaseP(cnt, price int) { // 포인터와 포인터 아닌 것으로 구분하면 됨
	b.cnt += cnt
	b.price += price
}

//원본 수정X (값복사 형식)
func (b shopingBasket) rePurchaseD(cnt, price int) {
	b.cnt += cnt
	b.price += price
}

func main() {
	//리시버 전달(값, 참조) 형식 [아~ 리시버는 메소드 처럼 작동하는구나!]
	//함수는 기본적으로 값 호출  -> 변수의 값이 복사 후 내부 전달(원본 수정X) -> 맵, 슬라이스 등은 참조 전달
	//리시버(구조체)도 마찬가지로 포인터를 활용해서 메소드 내에서 원본 수정 가능

	//예제1
	bs1 := shopingBasket{3, 5000}
	fmt.Println("ex1(totPrice) : ", bs1.purchase())
	bs1.rePurchaseP(10, 10000) //매개변수 전달(참조) : 구조체의 메소드에서 참조 형식으로 처리
	fmt.Println("ex1(totPrice) :", bs1.purchase())

	fmt.Println()

	//예제2
	bs2 := shopingBasket{5, 5000}
	fmt.Println("ex2(totPrice) : ", bs2.purchase())
	bs2.rePurchaseD(10, 10000) //매개변수 전달(복사)
	fmt.Println("ex2(totPrice) :", bs2.purchase())

	fmt.Println()

	//예제3
	bs3 := shopingBasket{10, 10000}

	fmt.Println("ex3(totPrice) : ", bs3.purchase())
	bs3.rePurchaseP(-5, -7000) //매개변수 전달(참조)
	fmt.Println("ex3(totPrice) :", bs3.purchase())
}
