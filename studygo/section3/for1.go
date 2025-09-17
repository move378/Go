package main

import "fmt"

func main() {
	/*
		고에서는 반복문이 for 만 있습니다. 반복문은 프로그래밍 언어의 특징이 담겨있는데
		예를 들면 파이썬에서는 generator, iterator, Lambda식을 통해서
		list나 dictionary 이런 자료구조를 빠르게 순회하기 때문에 파이썬을 많이 쓰는데,
		보통은 반복문을 통해서 데이터 전처리를 하게 됩니다. 형태소 분석을 하고, 머신러닝 등

		go 에서 유일한 반복문 for 그렇기 때문에 다양한 사용법을 숙지해야 합니다.
	*/

	//예제1
	for i := 0; i < 5; i++ {
		fmt.Println("ex1 : ", i)
	}

	/* 에러 발생
	   1) 행을 바꿨을 때, 엔터는 ; 이다.
	   2) 한 줄이라도 {}가 있어야 합니다.
	   3)
	*/

	//예제2
	// for {
	// 	fmt.Println("무한루프 ")
	// }

	//예제3 (Range 용법 : for index, value := range collection)
	loc := []string{"Seoul", "Busan", "Incheon"}
	for index, name := range loc { // index 를 _처리해서 값만 가져올 수 있습니다.
		fmt.Println("ex 3: ", index, name)
	}

	//예제4 (for문의 다양한 용법)
	sum1 := 0
	for i := 0; i <= 100; i++ {
		sum1 += i
	}
	fmt.Println("ex1 : ", sum1)

	sum2, i := 0, 0 // 가독성이 조금 더 좋아지는 형태
	for i <= 100 {
		sum2 += i
		i++
		// j := i++ // Go에서 후처리 연산은 반환 값이 없어서 컴파일 에러
	}
	fmt.Println("ex2 : ", sum2)

	//예제5 while문 형태 만들기
	sum3, i := 0, 0

	for { // while 형태와 비슷
		if i > 100 {
			break
		}
		sum3 += i
		i++
	}
	fmt.Println("ex3 : ", sum3)

	//예제4 변수가 두개
	for i, j := 0, 0; i <= 10; i, j = i+1, j+10 {
		fmt.Println("ex4 : ", i, j)
	}

	//에러 발생 for i, j := 0, 0; i <= 10 ; i++, j += 10 {

	/*
		❌ 에러 이유

		초기화 부분에서 여러 변수를 선언한 것 자체는 괜찮아요.
		i, j := 0, 0 → 가능 (Go에서 다중 선언 지원)

		후처리 부분에서 여러 문을 쓴 게 문제예요.
		i++, j += 10 → Go에서는 콤마(,)로 두 문장을 나열할 수 없습니다.
		후처리 부분에는 **하나의 단일 문(statement)**만 들어갈 수 있어요.

		즉, Go에서는 C처럼 for (i=0, j=0; ...; i++, j+=10) 문법이 안 됩니다.

		🔑 "후처리 연산에는 반환값이 없다"의 의미

		Go에서 **후처리 연산(post statement)**은 아래와 같은 특징이 있습니다:

		후처리 연산은 단일 statement
		즉, 하나의 문장만 들어갈 수 있고, 그 문장은 "값을 반환하지 않는 표현식"이어야 합니다.

		값을 반환할 수 없음
		Go의 후처리 부분은 루프 흐름 제어를 위한 side-effect(부수효과)만 허용합니다.
		예를 들어 i++, j++, callFunc() 처럼 결과값을 쓰지 않는 "statement"는 가능하지만,
		foo()처럼 값을 반환하는 함수를 써놓고 그 값을 어디에도 할당하지 않으면 컴파일 에러가 납니다.

		이유: 루프 제어를 예측 가능하게 유지
		Go는 불필요한 값을 무시하는 것을 허용하지 않으려 해서,
		후처리 부분에 반환값이 있는 표현식이 오면 반드시 변수에 할당하거나 사용해야 한다고 요구합니다.
	*/

	//예제 5 루프문의 흐름제어 (루프 레이블 사용)
Loop1: // 루프 레이블 밑에 관련이 없는 소스 코드가 있으면 컴파일 에러, 보통 반복문이나 조건문이 와야 합니다.
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 4 {
				break Loop1
			}
			fmt.Println("ex 5 : ", i, j)
		}
	}
	//예제 6 continue
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("ex6 : ", i)
	}
}
