//고루틴(Goroutine)기초(1)
package main

import "fmt"
import "time"

func exe1() {
	fmt.Println("exe1 func start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe1 func end", time.Now())
}

func exe2() {
	fmt.Println("exe2 func start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe2 func end", time.Now())
}

func exe3() {
	fmt.Println("exe3 func start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe3 func end", time.Now())
}

func main() {
	/* 고루틴(Goroutine)
	타 언어의 쓰레드(Thread:별도의 실행흐름)와 비슷한 기능
	생성 방법 매우 간단, 리소스 매우 적게 사용
	즉, 수많은 고루틴 동시 생성 실행 가능
	비동기적 함수 루틴 실행(매우 적은 용량 차지) -> 채널을 통해 통신
	공유메모리 사용 시에 정확한 동기화 코딩 필요
	싱글루틴에 비해 항상 빠른 처리 결과는 아니다.

	멀티 쓰레드의 장점과 단점
	장점 : 응답성 향상(웹서버), 자원공유를 효율적으로 활용 사용, 작업이 분리되어 코드 간결
	단점 : 구현하기 어려움, 테스트 및 디버깅 어려움, 전체프로세스의 사이드 이펙트, 성능 저하,
			동기화 코딩 반드시 숙지, 데드락... (파이썬 쓰레드를 배워볼 것 추천!)
	*/
	exe1() //가정 먼저 실행(일반적인 실행 흐름)

	//예제1
	fmt.Println("Main Routine Start : ", time.Now())
	go exe2()
	go exe3()
	time.Sleep(4 * time.Second) // 4초 대기 (주석처리하면 별도 고루틴 종료 : 메인함수 종료 시 모두 종료)
	//fmt.Scanln() //콘솔에서 테스트 할 경우 시간 대기 용도 사용 가능
	fmt.Println("Main Routine End : ", time.Now())
}
