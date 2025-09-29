/* 채널(Channel) 기초(3 Unbuffered Channel)

송신자가 ch <- true를 하면, 누군가 받을 때까지 블로킹
수신자가 <-ch를 하면, 누군가 보낼 때까지 블로킹

이로 인해 고루틴과 메인 함수가 번갈아가며 실행됨
채널의 동기화 메커니즘은 수신과 송신이 정확히 일치해야 함을 의미

unbuffered 채널의 동기적 특성은 데이터 교환의 정확한 순간적 동기화를 보장합니다.
송신자와 수신자가 정확히 만나는 순간에만 데이터 전달이 이루어지며, 이는 고루틴 간의 완벽한 협업 메커니즘을 제공합니다. */

package main

import (
	"fmt"
)

func main() {
	ch := make(chan bool)
	// ⭐ unbuffered channel 생성
	// 핵심: 송신자와 수신자가 "랑데뷰(만남)"해야만 데이터 전달 가능

	cnt := 6

	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true
			// ⚠️ 블로킹 포인트 1: 메인 고루틴이 <-ch로 받을 때까지 여기서 대기
			// 언블로킹 조건: 메인 고루틴이 <-ch 실행

			fmt.Println("Go : ", i)
			// 💡 이 출력은 메인 고루틴이 값을 받은 "직후"에 실행됨
		}
	}()

	for i := 0; i < cnt; i++ {
		<-ch
		// ⚠️ 블로킹 포인트 2: 고루틴이 ch <- true로 보낼 때까지 여기서 대기
		// 언블로킹 조건: 고루틴이 ch <- true 실행

		fmt.Println("Main : ", i)
		// 💡 이 출력은 고루틴이 값을 보낸 "직후"에 실행됨
	}

	// 📌 결과: 완벽한 핑퐁(ping-pong) 패턴
	// Go → Main → Go → Main → ...
}
