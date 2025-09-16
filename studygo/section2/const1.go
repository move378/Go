/* const 상수! final! 불변!
이 값은 정해진 최종값 입니다.

const 사용 초기화, 한번 선언되면 값 변경 금지, 고정된 값 관리용
*/

package main

import "fmt"

func main() {

	const a string = "Test1"

	/*
		상수는 선언과 동시에 초기화가 바로 이뤄져야 합니다.
		const g string
		g = "Test3"
		에러
	*/

	const b = "Test2"
	const c int32 = 10 * 10
	// const d = getHeight() 상수에서는 함수 변화되는 값 사용 불가
	const e = 35.6
	const f = false

	fmt.Println("a : ", a, "b : ", b, "c : ", c, "e : ", e, "f : ", f)
}
