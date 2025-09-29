/* 고루틴 동기화 기초(1 동기화 처리를 위한 객체 mutex 뮤텍스)
서로의 결과물을 머지한다든지 참조 메모리에 접근해서 데이터값을 수정한다던지
쓰고 지우고, 참조하고 뭐 등의 작업들은 동기화가 필수, 예상하지 않았던 결과물이 나올 수 있음
멀티 쓰레드 동기화

*/

package main

import (
	"fmt"
	"runtime" // mutax sync
)

// 구조체 선언(공유 데이터)
type count struct {
	num int
}

func (c *count) increment() {
	c.num += 1
}

func (c *count) result() {
	fmt.Println(c.num)
}

func main() {
	/* 고루틴 동기화 예제
	실행 흐름 제어 및 변수 동기화 가능
	공유 데이터 보호가 가장 중요 (Thread Safe)
	비즈니스 로직과 성능에 큰 영향 Java에서의 Synchronize
	다른 언어들과 달리 매우 빠르게 동기화 처리가 가능합니다.

	동기화 사용하지 않은 경우 예제
	시스템 전체 cpu 사용 */
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := count{num: 0}
	done := make(chan bool)

	for i := 1; i <= 10000; i++ {
		go func() {
			c.increment()
			done <- true
			runtime.Gosched() //Cpu 양보
		}()
	}

	for i := 1; i <= 10000; i++ {
		<-done
	}

	c.result()
}
