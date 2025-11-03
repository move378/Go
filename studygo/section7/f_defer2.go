// 함수 Defer(디퍼를 중첩되게 사용하는 경우에서...)
package main

import "fmt"

func start(t string) string {
	fmt.Println("start:", t)
	return t
}
func end(t string) {
	fmt.Println("end:", t)
}

func a() {
	defer end(start("b"))
	// 중첩되었을 때는 아래로 읽어내리는 흐름에 따라 함수 안의 함수나 값을 미리 실행해 놓고 결과나 값을 반환받은 상태에서
	// 밖에 있는 end 함수를 실행, 항상 중첩되어 있는 안에 있는 것부터 먼저 실행됨
	// defer에서 함수 중첩되게 쓰는 것 주의
	fmt.Println("in a")
}

func main() {

	//예제1
	a()
}
