//채널(Channel) 심화(2 함수에서 채널을 리턴하기)

package main

import (
	"fmt"
)

func sum(cnt int) <-chan int { // 리턴타입은 수신임(수신하는 놈이 받으면 됨)
	sum := 0
	tot := make(chan int)
	go func() {
		for i := 0; i <= cnt; i++ {
			sum += i
		}
		tot <- sum
	}() //익명 함수 클로져임. 이런 즉시실행 함수는 많이 사용함.
	return tot
}

func main() {
	//채널(Channel)
	//채널 또한 함수의 반환 값으로 사용 가능

	//예제1
	c := sum(100)

	fmt.Println("ex1 : ", <-c) // 동기식 여기서 수신

}
