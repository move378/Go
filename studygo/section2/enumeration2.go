//열거형2

package main

import "fmt"

func main() {

	const (
		A = iota
		// 규칙성을 가지고 증가되는 어떤 상수의 묶음을 만들 때
		B
		C
	)

	const (
		Jan = iota + 1
		Feb
		Mar
		Apr
		May
		Jun
	)

	fmt.Println(Jan)
	fmt.Println(Feb)
	fmt.Println(Mar)
	fmt.Println(Apr)
	fmt.Println(May)
	fmt.Println(Jun)

	fmt.Println(A, B, C)
}
