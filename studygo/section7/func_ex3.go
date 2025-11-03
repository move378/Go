// 함수 심화(재귀 함수- 스스로를 호출해서 어떤 프로세스의 결과값을 반환)
// [알고리즘에서 팩토리얼 계산 때 반드시 사용:팩토리얼이 뭔지는 아래 참고]
package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	} // 만약 매개변수 n의 인수로 0 을 받게 되면 끝나는 재귀함수
	return n * fact(n-1)
}

func prtHello(n int) {
	if n == 0 {
		return
	} // n이 0이 될때까지 실행하니까 n을 0으로 만드는 로직이 존재해야 함
	fmt.Println("ex2 : (", n, ")", "hi!")
	prtHello(n - 1)
}

func main() {
	/*	함수 고급
		재귀 함수(Recursion)
		프로그램이 보기 쉽고, 코드 간결, 오류 수정이 용이 : 장점
		코드 이해하기 어렵고, 기억공간을 코딩하는 능력치이 높아지면 많이 사용,
		제대로 끝내지 않으면 무한 루프 가능성[정확히 반환 및 끝나는 지점을 잘 생각해야 함] */

	//예제1
	x := fact(7) // 팩토리얼은 x * x-1, x * x-2,... 마지막에 1을 곱하고 끝나는 연산
	fmt.Println("ex1 :", x)

	//예제2
	prtHello(10)
}
