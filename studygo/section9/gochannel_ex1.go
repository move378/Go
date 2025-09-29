/* 채널(Channel) 심화(1 수신 및 송발신 전용 채널 : '매개변수'에 수신 및 발신 전용 채널 지정 가능)
수신 발신 전용 채널을 만들어보고, Select 문을 이용해서 Select case defualt
채널의 종류에 따라 혹은 채널에서 전송되는 값에 따라 다양한 분기 처리를 하는 활용방법
*/

package main

import (
	"fmt"
	"time"
)

func sendOnly(c chan<- int, cnt int) {
	// 송발신 전용 chan 키워드를 중심으로 화살표가 뒤 오른쪽 있으면 발신전용으로 매개변수를 사용하겠다.
	// 화살표를 보면 int형의 데이터를 채널로 담는 송발신
	for i := 0; i < cnt; i++ {
		c <- i
	}

	c <- 777

	//fmt.Println(<-c) //샌드 보내는 전용 채널에서 받는 것 수신을 처리 시 예외 발생
}

func receiveOnly(c <-chan int) {
	// 수신 전용 chan 키워드를 중심으로 화살표가 앞 왼쪽 있으면 수신전용으로 매개변수를 사용하겠다.
	// 화살표 모양을 보면 채널 int형으로부터 c로 값이 오는 [수신]
	for i := range c {
		fmt.Println("received : ", i)
	}

	fmt.Println(<-c) // 값을 수신 받아서 출력
}

func main() {
	/* 채널(Channel)
	함수 등의 매개변수에 수신 및 발신 전용 채널 지정 가능
	전용 채널 설정 후 방향이 다를 경우 예외 발생
	발신 전용 channel <- 데이터형
	수신 전용 <- channel
	매개 변수를 통해서 전용 채널 확인할 수 있다.

	채널 또한 함수의 반환 값으로 사용 가능 [채널을 리턴할 수 있음]
	*/

	//예제1
	c := make(chan int)

	go sendOnly(c, 10) //송발신전용
	go receiveOnly(c)  //수신전용

	time.Sleep(2 * time.Second) //2초간 대기

}
