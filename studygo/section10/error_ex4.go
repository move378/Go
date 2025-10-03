// Go 에러 처리 고급 (3. 예외 처리 구조체)

package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

// ============================================================
// PowError: 사용자 정의 에러 구조체
// ============================================================
// [덕타이핑 원리]
// Error() string 메소드만 구현하면, 명시적 선언(implements) 없이도
// error 인터페이스를 자동으로 구현한 것으로 인식됨
//
// Go는 상속이 없고 인터페이스만 존재하므로:
// - @override 같은 어노테이션 불필요
// - implements 키워드도 불필요
// - 메소드 시그니처만 일치하면 자동으로 인터페이스 구현 완료
type PowError struct {
	time    time.Time   // 에러 발생 시각
	value   interface{} // 에러를 발생시킨 파라미터 값
	message string      // 에러 설명 메시지
}

// ============================================================
// Error() 메소드 구현 - 덕타이핑의 핵심
// ============================================================
// 이 메소드를 구현하는 순간, PowError는 error 인터페이스를 만족함
//
// error 인터페이스 정의 (Go 표준 라이브러리):
//
//	type error interface {
//	    Error() string
//	}
//
// 리시버(Receiver): (e PowError)
// - 이 메소드를 PowError 타입에 연결
// - e는 PowError 인스턴스를 가리키는 변수명
func (e PowError) Error() string {
	return fmt.Sprintf("[%v]Error - Input Value(value: %g) - %s",
		e.time, e.value, e.message)
}

// ============================================================
// Power: 제곱 계산 함수 (에러 반환 포함)
// ============================================================
// 반환 타입: (float64, error)
// - Go의 관례: (결과값, 에러) 순서로 반환
//
// [다형성 적용]
// PowError 타입을 error 타입으로 반환 가능
// 이유: PowError가 Error() 메소드를 구현했으므로 (덕타이핑)
func Power(f, i float64) (float64, error) {
	// 잘못된 입력값에 대해 PowError 반환
	// PowError → error 타입으로 자동 변환 (다형성)
	if f == 0 {
		return 0, PowError{
			time:    time.Now(),
			value:   f,
			message: "0은 사용할 수 없습니다.",
		}
	}
	if math.IsNaN(f) {
		return 0, PowError{
			time:    time.Now(),
			value:   f,
			message: "첫 번째 파라미터가 숫자가 아닙니다.",
		}
	}
	if math.IsNaN(i) {
		return 0, PowError{
			time:    time.Now(),
			value:   i,
			message: "두 번째 파라미터가 숫자가 아닙니다.",
		}
	}

	// 정상 계산: 에러 없음 (nil 반환)
	return math.Pow(f, i), nil
}

func main() {
	// ============================================================
	// 예제1: 정상 케이스 (에러 없음)
	// ============================================================
	v, err := Power(10, 3)
	if err != nil {
		// [log.Fatal의 동작 원리]
		// 1. err.Error() 메소드 호출 → 에러 메시지 출력
		// 2. os.Exit(1) 호출 → 프로그램 즉시 종료
		//
		// [다형성 적용]
		// err의 선언 타입: error (인터페이스)
		// err의 실제 타입: PowError (구조체)
		// → 다형성 덕분에 err.Error() 호출 시 PowError.Error() 실행됨
		//
		// ⚠️ 주의: Fatal 이후의 코드는 실행되지 않음 (defer도 미실행)
		log.Fatal(err)
	}

	fmt.Println("ex1 : ", v) // 출력: ex1 : 1000

	// ============================================================
	// 예제2: 에러 발생 케이스 (f = 0)
	// ============================================================
	t, err := Power(0, 3)
	if err != nil {
		// [타입 어써션 (Type Assertion)]
		// 문법: 인터페이스변수.(구체타입)
		// 역할: error 인터페이스 → PowError 구체 타입으로 변환
		//
		// 왜 필요한가?
		// - error 인터페이스로는 Error() 메소드만 접근 가능
		// - PowError의 추가 필드(time, value, message)에는 접근 불가
		// - 구체 타입으로 변환해야 추가 필드 접근 가능
		//
		// [비유]
		// 다형성: PowError를 error "택배 상자"에 포장
		// 타입 어써션: error "택배 상자"를 열어서 PowError 꺼내기
		//
		// ⚠️ 주의: 아래 방식은 위험함 (타입이 다르면 panic 발생)
		// ✅ 안전한 방식은 밑에서 설명
		fmt.Println(err.(PowError).value)   // 입력값 출력
		fmt.Println(err.(PowError).message) // 에러 메시지 출력
		fmt.Println(err.(PowError).time)    // 에러 시각 출력
	}

	fmt.Println("ex2 : ", t) // 출력: ex2 : 0

	// ============================================================
	// 예제3: 안전한 타입 어써션 (권장 방식)
	// ============================================================
	fmt.Println("\n=== 안전한 타입 어써션 예제 ===")
	_, err = Power(math.NaN(), 2)
	if err != nil {
		// [안전한 타입 어써션 문법]
		// if 변수, ok := 인터페이스변수.(구체타입); ok { ... }
		//
		// ok == true:  변환 성공, 변수에 구체 타입 값 저장됨
		// ok == false: 변환 실패, 다른 타입이었음 (panic 발생 안 함)
		//
		// 왜 안전한가?
		// - 어떤 타입이 들어있을지 장담할 수 없는 상황에서
		// - if문으로 타입을 체크하고 안전하게 사용 가능
		if powErr, ok := err.(PowError); ok {
			fmt.Printf("✅ PowError 확인됨\n")
			fmt.Printf("  - 발생 시각: %v\n", powErr.time.Format("15:04:05"))
			fmt.Printf("  - 입력 값: %v\n", powErr.value)
			fmt.Printf("  - 메시지: %s\n", powErr.message)
		} else {
			fmt.Println("❌ PowError가 아닌 다른 에러 타입")
		}
	}

	// ============================================================
	// 예제4: 타입 스위치 (여러 에러 타입 처리)
	// ============================================================
	fmt.Println("\n=== 타입 스위치 예제 ===")
	errors := []error{
		PowError{time: time.Now(), value: 0, message: "0 입력"},
		fmt.Errorf("일반 에러"),
	}

	for i, e := range errors {
		fmt.Printf("에러 %d: ", i+1)

		// switch문으로 타입별 처리
		switch specificErr := e.(type) {
		case PowError:
			fmt.Printf("PowError - %s (값: %v)\n",
				specificErr.message, specificErr.value)
		default:
			fmt.Printf("기타 에러 - %v\n", specificErr)
		}
	}
}

// ============================================================
// 핵심 개념 정리
// ============================================================
//
// 1. [덕타이핑 (Duck Typing)]
//    - Error() string 메소드만 구현하면
//    - implements 키워드 없이도 error 인터페이스 자동 구현
//    - Go의 암묵적 인터페이스 구현 방식
//
// 2. [다형성 (Polymorphism)]
//    - 구체적 타입(PowError) → 추상적 타입(error)으로 추상화
//    - 여러 에러 타입을 하나의 error 타입으로 통일해서 처리
//    - 코드 재사용성과 유연성 증가
//
// 3. [타입 어써션 (Type Assertion)]
//    - 추상적 타입(error) → 구체적 타입(PowError)으로 구체화
//    - 인터페이스에서 접근 불가능한 추가 필드/메소드 접근 가능
//    - 안전한 사용을 위해 if _, ok := ... 패턴 권장
//
// 4. [log.Fatal]
//    - err.Error() 호출 → 에러 메시지 출력
//    - os.Exit(1) 호출 → 프로그램 즉시 종료
//    - defer 함수도 실행되지 않으므로 주의 필요
//
// ============================================================
