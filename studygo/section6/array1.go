// 자료형 : 배열(1)
package main

import "fmt"

func main() {

	/* 배열
	배열은 용량, 길이 항상 같다.
	배열 vs 슬라이스 차이점 중요
	길이 고정 						   vs 길이 가변
	값 타입 								  vs 참조 타입
	복사 전달 							 vs  참조 값 전달
	전체 비교연산자 사용 가능 vs 비교 연산자 사용 불가
	대 부분 슬라이스 사용한다.

	cap() : 배열, 슬라이스 용량
	len() : 배열, 슬라이스 개수
	*/

	//예제1
	var arr1 [5]int
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5}
	var arr3 = [5]int{1, 2, 3, 4, 5}
	arr4 := [5]int{1, 2, 3, 4, 5}
	arr5 := [5]int{1, 2, 3} //기본 0 초기화
	arr6 := [...]int{1, 2, 3, 4, 5}
	arr7 := [5][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10}, //콤마 주의
	}

	arr1[2] = 5 //값 삽입

	fmt.Printf("%-5T %d %v\n", arr1, len(arr1), arr1)
	fmt.Printf("%-5T %d %v\n", arr2, len(arr2), arr2)
	fmt.Printf("%-5T %d %v\n", arr3, len(arr3), arr3)
	fmt.Printf("%-5T %d %v\n", arr4, len(arr4), arr4)
	fmt.Printf("%-5T %d %v\n", arr5, len(arr5), arr5)
	fmt.Printf("%-5T %d %v\n", arr6, len(arr6), arr6)
	fmt.Printf("%-5T %d %v\n", arr7, len(arr7), arr7)
	// %dex, %value, %-5길이Type printf 함수에서는 몇가지 조정 옵션이 있습니다.

	//예제2
	arr8 := [5]int{1, 2, 3, 4, 5}
	arr9 := [5]int{ //여러 줄 선언 콤마 주의
		1,
		2,
		3,
		4,
		5,
	}
	arr10 := [...]string{"Kim", "Lee", "Park"}

	fmt.Printf("%-5T %d %v\n", arr8, len(arr8), arr8)
	fmt.Printf("%-5T %d %v\n", arr9, len(arr9), arr9)
	fmt.Printf("%-5T %d %v\n", arr10, len(arr10), arr10)

	//배열 순회
	//예제3
	arr11 := [5]int{1, 10, 100, 1000, 10000}

	//len 길이 반복
	for i := 0; i < len(arr11); i++ {
		fmt.Println("ex1 : ", arr11[i])
	}
	fmt.Println() //줄 바꿈

	//예제4
	arr12 := [5]int{1, 10, 100, 1000, 10000}

	//range 사용
	for i, v := range arr12 {
		fmt.Println("ex2 : ", i, v)
	}
	fmt.Println()

	//인덱스 생략1
	for _, v := range arr12 {
		fmt.Println("ex3 : ", v)
	}

	//인덱스 생략2
	fmt.Println()
	for v := range arr12 { // 첫번째 인자는 인덱스 i, 두번째 인자가 값임.
		fmt.Println("ex4 : ", v)
	}

	//배열 복사
	//값 복사 확인 중요

	//예제5
	arr13 := [5]int{1, 10, 100, 1000, 10000}
	arr14 := arr13 // 여기서 얕은 복사가 아니라 그냥 깊은 복사가 됨, 언어별 차이!
	// 왜냐하면 go에서 arry는 참조타입이 아니라 값 타입이다!! 그렇기 때문에 깊은 복사가 이뤄진다.

	fmt.Println("ex1 : ", arr13, &arr13)
	fmt.Println("ex1 : ", arr14, &arr14)

	fmt.Printf("ex1: %p %v\n", &arr13, arr13) //주소 값 출력
	fmt.Printf("ex1: %p %v\n", &arr14, arr14) //주소 값 출력
	//&를 붙이면 주소에 있는 값을 출력해줍니다.
	//%p 는 포인터, %v 는 오리지널 밸류, %p 에는 주소값 &로 매핑해줘야 합니다.

	arr13[2] = 200
	fmt.Printf("ex1: %p %v\n", &arr13, arr13) // arr13[2] 200
	fmt.Printf("ex1: %p %v\n", &arr14, arr14) // 깊은 복사 확인

}
