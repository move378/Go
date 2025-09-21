// 자료형 : 슬라이스(2)
package main

import "fmt"

func main() {
	//슬라이스(슬라이스 참조 타입 증명)

	//예제1(배열:복사)
	arr1 := [3]int{1, 2, 3}
	var arr2 [3]int

	arr2 = arr1
	arr2[0] = 7

	fmt.Println("ex1 : ", arr1)
	fmt.Println("ex1 : ", arr2)
	fmt.Println()

	//예제2(슬라이스:얕은 복사)
	slice1 := []int{1, 2, 3}
	var slice2 []int

	slice2 = slice1
	slice2[0] = 7

	fmt.Println("ex2 : ", slice1)
	fmt.Println("ex2 : ", slice2)
	fmt.Println()
	/*
	대용량 데이터를 매개변수로 넘길 때, 원본이 수정이 된다는 것은 5개의 변수에 하나의 슬라이스를 할당을 해도
	원본 값이 같이 수정이 되기 때문에 주로 메소드의 파라미터로 넘길 때 사용을 많이 합니다.
	슬라이스를 사용하는 이유가 여기에 있습니다(얕은 복사:참조 복사)
	*/

	//예제3(슬라이스 예외 상황)
	slice3 := make([]int, 5, 10) //여기서 용량만큼 초기화 되는게 아니라 길이만큼 초기화 됩니다.
	fmt.Println("ex3 : ", slice3[4])
	//fmt.Println("ex3 : ", slice3[5]) //길이 index over 예외
	//fmt.Println("ex3 : ", slice3[8]) //길이 index over 예외
	fmt.Println()

	//예제4(슬라이스 순회)
	for i, v := range slice1 {
		fmt.Println("ex4 : ", i, v)
	}
}
