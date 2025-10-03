//Go 에러 처리 고급(3 예외 처리 구조체)

package main

import (
	"fmt"
	"log" // fatal 때문에 log가 있어야 함.
	"math"
	"time"
)

//예외(에러) 처리 구조체, 덕타입-인터페이스에 있는 메소드를 구현하면 오리나 개가 되는 것임. 짓다 물다
//즉, Error() 메소드를 구현한다면 Error를 담당하는 구조체가 됩니다 - 다형성
type PowError struct {
	time    time.Time //에러 발생 시간
	value   interface{}   //파라미터
	message string    //에러 메시지
}

func (e PowError) Error() string {
	// 리시버로 해당 구조체를 받아서 연결하기, Error 메소드를 구현![덕타이핑] - Error의 리턴타입은 스트링
	return fmt.Sprintf("[%v]Error - Input Value(value: %g) - %s", e.time, e.value, e.message)
}

func Power(f, i float64) (float64, error) {
	// Error() 메소드를 구현하였기 때문에 PowError을 error처럼 사용할 수 있는 원리로 Error()의 반환값 String이 출력됨
	if f == 0 {
		return 0, PowError{time: time.Now(), value: f, message: "0은 사용할 수 없습니다."}
	}
	if math.IsNaN(f) {
		return 0, PowError{time: time.Now(), value: f, message: "숫자가 아닙니다."}
	}
	if math.IsNaN(i) {
		return 0, PowError{time: time.Now(), value: i, message: "숫자가 아닙니다."}
	}
	return math.Pow(f, i), nil
}

func main() {
	//에러 처리 고급
	//error 타입이 아닌 경우 에러 처리 방법
	//Error 메소드를 구현해서 사용자 정의 에러 처리 예제 심화
	//구조체를 사용해서 세부적인 정보 출력[기존 IDE가 해주는 기능]

	//예제1
	v, err := Power(10, 3) //정상 처리
	if err != nil {
		log.Fatal(err) // 다형성의 원리로써 error 구조체를 인식함 error 객체 override
		//log.Fatal(err.Error())
	}

	fmt.Println("ex1 : ", v)

	//예제2
	t, err := Power(0, 3) //예외 발생
	if err != nil {
		//log.Fatal(err)
		//log.Fatal(err.Error())

		//타입 어써션에 대한 개념, 타입 변환
		fmt.Println(err.(PowError).value)
		fmt.Println(err.(PowError).message)
		fmt.Println(err.(PowError).time)
	}

	fmt.Println("ex2 : ", t)

}
