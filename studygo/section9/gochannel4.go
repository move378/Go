//채널(Channel) 기초(4 비동기식 : use buffered channel - 고의 병행성을 구현)

package main

import (
	"fmt"
	"runtime"
)

func main() {
	/* 채널(Channel)
	예제1(비동기 : 버퍼 사용)

	버퍼
	발신 -> 가득차면 대기, 비어있으면 작동,
	수신 -> 비어있으면 대기, 가득차있으면 작동 */

	runtime.GOMAXPROCS(1)
	// 멀티 코어에서의 비동기 채널을 사용할 때에는실행 순서나 채널 사용 방법에 대해 신경을 써야 합니다.
	// ex) 여러 프로세서가 처리하기 때문에 순서가 뒤바뀜
	ch := make(chan bool, 4)
	cnt := 12

	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true
			fmt.Println("Go : ", i)
		}
	}()

	for i := 0; i < cnt; i++ {
		<-ch
		fmt.Println("Main : ", i)
	}

}
