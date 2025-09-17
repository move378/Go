package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//예제1 golang 에서 스위치 문은 특별하다 시간될 때 레퍼런스 공식문서를 찾아 볼 것!
	rand.Seed(time.Now().UnixNano())
	switch i := rand.Intn(100); {
	case i >= 50 && i < 100:
		fmt.Println("i -> ", i, " 50 이상 100 미만")
	case i >= 25 && i < 50:
		fmt.Println("i -> ", i, " 25 이상 50 미만")
	default:
		fmt.Println("i -> ", i, " 기본 값")
	}

	//예제 2 여러 가지 값을 나열해서 케이스 문에 걸리는지를 확인
	switch a := 30 / 15; a {
	case 2, 4, 6: // a가 2, 4, 6인 경우
		fmt.Println("a -> ", a, "는 짝수")
		if a == 2 {
			fmt.Println("a는 2")
		} // 이렇게 if문으로 다시 거를 수도 있음
	case 1, 3, 5: // a가 1, 3, 5인 경우
		fmt.Println("a -> ", a, "는 짝수")
	}

	//예제 3
	switch e := "go"; e {
	case "java":
		fmt.Println("Java!")
	case "go":
		fmt.Println("go!")
		//break go에서는 break 가 생략되어 있다. 그런데 fallthrough 가 있으면
		fallthrough // 예외처리, 요청에 대한 응답 같은 부분에서 간혹 사용됩니다.
	case "python":
		fmt.Println("python")
	case "ruby":
		fmt.Println("ruby")
	case "rust":
		fmt.Println("rust")

	}

}
