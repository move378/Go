// 함수 Closure(2)
package main

import "fmt"

func main() {

	//예제1
	cnt := increaseCnt()

	fmt.Println("ex1 : ", cnt())
	fmt.Println("ex1 : ", cnt())
	fmt.Println("ex1 : ", cnt())
	fmt.Println("ex1 : ", cnt())
	fmt.Println("ex1 : ", cnt())

}

func increaseCnt() func() int { //반환 형이 int인 함수를 리턴하는 increaseCnt()를 정의
	n := 0
	//지역변수(캡처되어 참조를 유지, 계속 0으로 초기화되는 것이 아니라 참조를 유지함!)
	//아래의 익명함수를 가진 함수가 변수에 할당 될 때, 캡쳐되어 클로져가 할당된 변수가 소멸 때까지 참조를 유지함!!
	// "클로저가 생성될 때의 스코프에 있는 변수들과 생명주기를 공유"
	return func() int {
		n += 1
		return n
	}
}
