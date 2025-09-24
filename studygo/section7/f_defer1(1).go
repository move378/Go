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
	defer end(start("b")) // 중첩되었을 때는 안에 있는 것을 실행시키고 밖에 있는 end 함수를 마지막에 실행
	fmt.Println("in a")
}

func main() {

	//예제1
	a()
}
