// 자료형 : 포인터(1)
/* Go에서 포인터를 어디까지 지원을 하는지? 다른 언어와 차이점은 무엇인지?
포인터를 알아야 구조체를 배울 때 레퍼런스 타입으로 함수에 넘길 수가 있는데,
포인터가 C에서 처럼 주소값에 대한 연산을 지원하지 않기 때문에 개념만 알면 되서 쉽습니다. 

*** [&X - X의 주소값, *주소값을가진변수 - 해당 주소가 가르키는 실제 값] ***
*/

package main

import "fmt"

func main() {
	/* 포인터
	Go : 포인터 지원(C, C++) [변수의 지역성, 연속된 메모리 참조 ... 힙, 스택....]
	Go에서는 지역 참조성에 대한 지원만 해주어
	주소의 값은 직접 변경 불가능(잘못된 코딩으로 인한 버그 방지)
	메모리 주소를 출력(값의 메모리 주소)
	애스터리스크로 사용(*)
	nil로 초기화(0 아님)
	*/

	//예제1
	var a *int            //방법1
	var b *int = new(int) //방법2 : int 크기의 공간이 지정된 주소를 할당함

	fmt.Println(a) // nil로 초기화되며, *(포인터형:주소 번지수)으로 &(해당하는 주소 번짓수)을 붙여서 주소 번짓수를 전달해야함
	fmt.Println(b) // *int

	i := 7
	a = &i // &i(i에 대한 주소 번지수) 전달
	b = &i // b도 똑같은 값을 전달

	var c = &i //방법3
	d := &i    //방법4

	fmt.Println("ex1 : ", i)  // 7
	fmt.Println("ex1 : ", &i) //실행 할 때 마다 시스템 별로 변경

	fmt.Println("ex1 : ", *a) //i의 주소값을 찾아가서 거기에 있는 실제 값 [역참조라함] 7
	fmt.Println("ex1 : ", a)  //i의 주소값이 담겨진 포인터 변수 a [주소 번지수]
	fmt.Println("ex1 : ", &a) //포인터 변수 a의 주소 값
	// 정리 &를 붙이면 무조건 주소 번지수를 반환하기에 &를 사용해야 하고, 주소 번지수에 * 붙이면 주소 번지수를 찾아가서 실제 값을 가져옴[역참조]

	fmt.Println("ex1 : ", *b)
	fmt.Println("ex1 : ", b)
	fmt.Println("ex1 : ", &b) //포인터 변수 b의 주소 값

	fmt.Println("ex1 : ", *c)
	fmt.Println("ex1 : ", c)
	fmt.Println("ex1 : ", &c) //포인터 변수 c의 주소 값

	fmt.Println("ex1 : ", *d)
	fmt.Println("ex1 : ", d)
	fmt.Println("ex1 : ", &d) //포인터 변수 d의 주소 값

	//예제2
	// i := 7
	p1 := &i

	fmt.Println("ex2 : ", i, *p1, &i, p1) // 이렇게 4가지만 이해하면 됩니다.

	*p1++ // go에서 주소값은 연산할 수 없지만, 역참조한 실제 값은 연산이 가능}
	fmt.Println("ex2 : ", i, *p1, &i, p1)

	*p1 = 10 //포인터 변수 역 참조 값 변경 - 문제 없음!
	fmt.Println("ex2 : ", i, *p1, &i, p1)

	fmt.Println("ex2 : ", i, *p1, &i, p1)

	//예제3
	type bag struct{ witdh, height, weight float32 } // 구조체
	var p *int = new(int)                            // p는 int형의 포인터 타입 (기본 타입)
	var p_bag *bag = &bag{20, 50, 30}                // p_bag는 bag 포인터 타입 (구조체)

	/*
		둘 다 포인터가 맞습니다 - 메모리 주소를 저장
		fmt.Println의 출력 방식이 다를 뿐

		*int → 주소 출력
		*구조체 → &{필드값들} 형태로 출력 (개발자 편의성)
	*/

	fmt.Println("ex3 : ", p)
	fmt.Println("ex3 : ", p_bag)
	fmt.Println()

	//p++              //컴파일 에러, 포인터 연산 허용 X
	//p = 0xc071405232 ////컴파일 에러, 주소값 대입 허용 X

}
