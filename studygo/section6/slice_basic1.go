// 자료형 : 슬라이스(1)

package main

import "fmt"

func main() {
	/*
		길이x(가변) -> 동적으로 크기가 늘어난다. , 레퍼런스(참조 값) 타입
		어떤 메소드로 값을 전달할 때 99% 퍼포먼스 성능 상의 이유로 슬라이스 타입을 매개변수로 전달합니다.
		슬라이스( 길이 & 용량 )크기가 동적으로 할당 가능

		배열 vs 슬라이스 차이점 [중요]
		길이 고정  			vs 길이 가변
		값 타입			 	vs 참조 타입
		복사 전달 			 vs  참조 값 전달
		전체 비교연산자 사용 가능 vs 비교 연산자 사용 불가
		대 부분 슬라이스 사용한다.

		cap() : 용량
		len() : 개수

		2가지 선언 방법 1) 배열처럼 선언, 2) make 함수 : make(자료형, 길이, 용량(생략시 길이))
	*/

	//예제1
	var slice1 []int
	slice2 := []int{}
	slice3 := []int{1, 2, 3, 4, 5}
	// Object-Relational Mapping 데이터의 로우를 객체로 가져와서 슬라이스에 담아서 키 : 값으로 가져오는 형식
	slice4 := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	// slice2[0] = 1
	/*
		길이가 확실한 경우에는 나머지를 0으로 자동 초기화하지만 슬라이스는 길이가 가변형으로 원소를 넣어놓지 않은 상태로
		특정 인덱스 값을 수정하면 런타임 범위에러가 발생합니다. 즉 범위가 지정되지 않을 경우에는 0으로 초기화 되어
		자동으로 메모리 공간 초기화해서 만들어 놓지 않습니다. 값이 초기화되지 않은 상태에서는 값을 수정할 수 없습니다.
		왜 가변형이기 때문에!
	*/

	slice3[4] = 10 // 이미 길이가 확정되어 해당 인덱스의 값을 수정가능
	fmt.Printf("%-5T %d %d %v\n", slice1, len(slice1), cap(slice1), slice1)
	fmt.Printf("%-5T %d %d %v\n", slice2, len(slice2), cap(slice2), slice2)
	fmt.Printf("%-5T %d %d %v\n", slice3, len(slice3), cap(slice3), slice3)
	fmt.Printf("%-5T %d %d %v\n", slice4, len(slice4), cap(slice4), slice4)

	// 예제2
	var slice5 []int = make([]int, 5, 10)
	/*
		slice 쌩으로 생성과 make 함수 이용시 비교 - make로 만들면 자료형에 따라 초기화가 자동으로 됩니다.
		그래서 make를 많이 사용합니다.

		(자료형, 길이, 용량(생략시 길이와 같이 맞춤))
		성능 최적화 부분에서 미리 사용할 용량에 근접하게 확보해 놓은 뒤에 데이터를 저장해 사용하는 것이 제일 좋습니다.
		이 부분이 늘어나면 메모리에 재할당이 일어나고, 확보한 용량보다 너무 적게 쓰면 메모리에 낭비가 발생합니다.
		go에서 잘 만들어진 부분은 이렇게 길이와 용량을 성능을 고려해서 좀 디테일 하게 설계가 가능하도록 만들어놨습니다.
		숙련된 개발자들은 이부분에 대한 고민이 많기 때문에 굉장히 좋아 합니다!
	*/

	var slice6 = make([]int, 5) // 이걸 가장많이 사용
	slice7 := make([]int, 5, 100)
	slice8 := make([]int, 5)

	slice6[2] = 7 // make로 생성시 초기화 된 상태이기 때문에 값 수정 가능
	fmt.Printf("%-5T %d %d %v\n", slice5, len(slice5), cap(slice5), slice5)
	fmt.Printf("%-5T %d %d %v\n", slice6, len(slice6), cap(slice6), slice6)
	fmt.Printf("%-5T %d %d %v\n", slice7, len(slice7), cap(slice7), slice7)
	fmt.Printf("%-5T %d %d %v\n", slice8, len(slice8), cap(slice8), slice8)

	//예제3
	var slice9 []int // nil 상태 : 슬라이스(길이와 용량 0) 값을 수정할 수 없습니다.

	if slice9 == nil {
		fmt.Println("This is Nil Slice!")
	}

}
