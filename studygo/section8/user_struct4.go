/* 사용자 정의 타입- 4 포인터 리시버와 값 리시버
값 리시버: 구조체의 복사본을 받음
포인터 리시버: 구조체의 메모리 주소를 받음
Go: 명시적 복사 vs 참조 선택 

메소드가 구조체를 수정하나요?
├─ YES → 포인터 리시버 (*T)
└─ NO
   ├─ 구조체가 큰가요? (>64바이트)
   │  ├─ YES → 포인터 리시버 (*T)
   │  └─ NO → 값 리시버 (T)
   └─ 동시성 안전이 중요한가요?
      ├─ YES → 값 리시버 (T)
      └─ NO → 포인터 리시버 (*T)
*/
package main

import "fmt"

// 타입정의 장바구니 구조체{ cnt와 price int 를 가지는 사용자 정의 구조체 }
type shopingBasket struct{ cnt, price int }

// 구매 함수
func (b shopingBasket) purchase() int {
	//^^^^^^^^^^^^^^^ 부분이 리시버
	return b.cnt * b.price
}

// 원본 수정 (포인터 리시버) - 포인터 리시버로 해당 구조체 타입을 넘기면 * 넣어서 역참조하고 그럴 필요 없음
func (b *shopingBasket) rePurchaseP(cnt, price int) { // 포인터와 포인터 아닌 것으로 구분하면 됨
	b.cnt += cnt
	b.price += price
}

// 원본 수정X (값 리시버)
func (b shopingBasket) rePurchaseD(cnt, price int) {
	b.cnt += cnt
	b.price += price
}

func main() {
	/* 리시버 전달(참조 전달) 형식
	함수는 기본적으로 값 호출로 -> 변수의 값이 복사 후 내부 전달(원본 수정X) [여기서도 맵, 슬라이스 등은 참조 전달]
	포인터 리시버로 참조전달을 할 수 있음. 즉, 원본 수정 가능 */

	//예제1
	bs1 := shopingBasket{3, 5000}
	fmt.Println("ex1(totPrice) : ", bs1.purchase())
	bs1.rePurchaseP(10, 10000) // 매개변수 전달 (참조)
	fmt.Println("ex1(totPrice) :", bs1.purchase())

	fmt.Println()

	//예제2
	bs2 := shopingBasket{5, 5000}
	fmt.Println("ex2(totPrice) : ", bs2.purchase())
	bs2.rePurchaseD(10, 10000) // 매개변수 전달(복사)
	fmt.Println("ex2(totPrice) :", bs2.purchase())

	fmt.Println()

	//예제3
	bs3 := shopingBasket{10, 10000}

	fmt.Println("ex3(totPrice) : ", bs3.purchase())
	bs3.rePurchaseP(-5, -7000) //매개변수 전달(참조)
	fmt.Println("ex3(totPrice) :", bs3.purchase())
}
