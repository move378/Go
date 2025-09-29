//채널(Channel) 심화(4 셀렉트 응용하여 분기처리하기)

package main

import (
	"fmt"
	"time"
)

func main() {
	//채널(Channel) 셀렉트 구문
	//채널 Select 구문 -> 채널에 값이 수신되면 해당 case 문 실행
	//일회성 구문이므로, For(반복문)안에서 수행
	//default 구문 처리 주의

	//예제1
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		for {
			ch1 <- 77
			time.Sleep(250 * time.Millisecond)
		}
	}()

	go func() {
		for {
			ch2 <- "Golang Hi!"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select { // ch1/ch2 어디서 오는지로 분기함
			case num := <-ch1:
				fmt.Println("ch1 : ", num)
			case str := <-ch2:
				fmt.Println("ch2 : ", str)
				//default:
				//fmt.Println("default test")
				//값을 받아와야 하는데, 값이 오기 전이라 계속 디폴트로 빠지는 반복문이 되어버림
			}
		}
	}()

	time.Sleep(7 * time.Second) // main 함수를 7초간 실행함 이거 없으면 무한반복
}
