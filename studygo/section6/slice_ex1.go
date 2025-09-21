package main

import (
	"fmt"
	"sort"
)

func main() {
	// 슬라이스 수정과 삽입 및 병합
	// 예제1

	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{8, 9, 10, 11, 12}
	s3 := []int{13, 14, 15, 16, 17}
	s1 = append(s1, 6, 7)
	s2 = append(s1, s2...)      //슬라이스를 삽입 할 경우 ... 사용
	s3 = append(s2, s3[0:3]...) //추출 후 병합

	/*
		길이를 5 용량을 10이라고 했을 때, 길이가 7, 용량이 10이 되는데, 길이가 11로 넘어가는 순간
		s1에 같은 용량이 10인 슬라이스를 초기화해서 값과 메모리를 추가 할당합니다.
		이부분이 대용량으로 넘어갈 때는 큰 스트레스 작용할 수 있는데,
		세밀하게 코딩하려면 용량까지 예측해서 지정해놔야 로딩 시간이 단축되고 성능에 관한 향상으로 이뤄집니다. 좀더 신경쓸 것
		내부적으로 용량을 벗어나면 용량 *2인 슬라이스가 내부적으로 배열로 옮겨졌다가 생성이 된다는 것을 강조함.
	*/
	fmt.Println("ex1 : ", s1)
	fmt.Println("ex1 : ", s2)
	fmt.Println("ex1 : ", s3)
	fmt.Println()

	//예제2
	s4 := make([]int, 0, 5)
	for i := 0; i < 15; i++ {
		s4 = append(s4, i)
		fmt.Printf("ex2 -> len: %d, cap: %d, value: %v\n", len(s4), cap(s4), s4)
		//길이 및 용량 자동 증가(용량 : 2배)
	}

	/* 슬라이스 추출 및 정렬
	슬라이스 부분은 파이썬에서 가져온 것!
	
	slice[i:j] s -> j-1 까지 추출
	slice[i:]  i -> 마지막 까지 추출
	slice[:j] 처음 -> j-1 까지 추출
	slice[:]  처음부터 마지막 까지 추출
	*/

	//예제3 (추출)
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("ex1 : ", slice1[:])
	fmt.Println("ex1 : ", slice1[0:])
	fmt.Println("ex1 : ", slice1[:5])
	fmt.Println("ex1 : ", slice1[0:len(slice1)])
	fmt.Println("ex1 : ", slice1[3:])
	fmt.Println("ex1 : ", slice1[:3])
	fmt.Println("ex1 : ", slice1[1:3])
	fmt.Println()

	//예제4 (정렬)
	//sort 패키지 : https://golang.org/pkg/sort/ 참조
	slice2 := []int{3, 6, 10, 9, 1, 4, 5, 8, 2, 7}
	slice3 := []string{"b", "d", "f", "a", "c", "e"}
	fmt.Println("ex2 : ", sort.IntsAreSorted(slice2))
	sort.Ints(slice2)
	fmt.Println("ex2 : ", slice2) //Int형 정렬
	fmt.Println("ex2 : ", sort.IntsAreSorted(slice2))

	fmt.Println()

	fmt.Println("ex2 : ", sort.StringsAreSorted(slice3))
	sort.Strings(slice3)
	fmt.Println("ex2 : ", slice3) //String형 정렬
	fmt.Println("ex2 : ", sort.StringsAreSorted(slice3))

	/*슬라이스 복사
	copy(복사 대상, 원본)
	make로 공간을 할당 후 복사 해야한다. [공간 할당 없으면 얕은 복사]
	make / new 뭐 이런 같은 류의 키워드임
	복사 된 슬라이스 값 변경해도 원본에는 영향 없음
	*/

	//예제1(복사)
	slice4 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice5 := make([]int, 5)
	slice6 := []int{}

	copy(slice5, slice4) // 5번의 용량이 5이기 때문에 5개만 복사
	copy(slice6, slice4) // 길이가 없어, make로 복사하지 않았기 때문에 얘는 공간 자체가 없는 얘임

	fmt.Println("ex1 : ", slice5)
	fmt.Println("ex1 : ", slice6) //복사 안됨

	fmt.Println()

	//예제2
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5)
	copy(b, a)
	b[0] = 7
	b[4] = 10

	fmt.Println("ex2 : ", a) //원본 유지
	fmt.Println("ex2 : ", b) //복사된 대상 변경
	fmt.Println()

	//예제3
	c := [5]int{1, 2, 3, 4, 5}
	d := c[0:3] //주의! 부분적으로 슬라이스 추출은 참조만 가져오는 것[얕은 복사] -> 원본 값 변경 된다.
	d[1] = 7    // d 의 경우 make로 새로 생성하지 않았어! 해당 인덱스의 주소만 가져오는 결과
	// 외워야 할 부분은 기초과정에서 외워야 합니다. 왜? 그 언어를 만든 사람이 규칙을 그렇게 만들었으니까

	fmt.Println("ex3 : ", c)
	fmt.Println("ex3 : ", d)
	fmt.Println()

	//예제4
	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f := e[0:5:7] //용량 지정, 마지막 인덱스가 용량을 뜻하는 용량지정임

	fmt.Println("ex4 : ", len(f), cap(f))
	fmt.Println("ex4 : ", f)

	f[0] = 7
	f[4] = 10

	fmt.Println("ex4 : ", e)
	fmt.Println("ex4 : ", len(f), cap(f))
	fmt.Println("ex4 : ", f) // 그렇지 주소만 가져온거야 근데 용량이 7이라는 건 신기하네

}
