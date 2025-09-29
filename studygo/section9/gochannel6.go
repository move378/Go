//채널(Channel) 기초(5 값, 결과 변수에 할당 패턴)

package main

import (
	"fmt"
)

func main() {
	//채널(Channel)
	//Close : 채널 닫기

	ch := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- 77777
		}
	}()

	val1, ok1 := <-ch // 두번째 변수에 전송 결과를 boolean 값으로 전달함
	fmt.Println("ex1 : ", val1, ok1)
	val2, ok2 := <-ch
	fmt.Println("ex1 : ", val2, ok2)
	val3, ok3 := <-ch
	fmt.Println("ex1 : ", val3, ok3)

	close(ch) //채널 닫기
	val4, ok4 := <-ch
	fmt.Println("ex1 : ", val4, ok4)
}
