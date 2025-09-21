//데이터 타입 : Numeric 연산(1)

package main

import (
	"fmt"
	"math"
)

func main() {
	//숫자 연산(산술, 비교)
	//타입이 같아야 가능
	//형변환 시 중요한 값은 계산기로 계산을 해볼 것
	//다른 타입끼리는 반드시 형 변환 후 연산
	//형 변환이 없을 시 예외(에러) 발생
	// +, -, *, %, /, <<, >>, $, ^ 반전

	// 예제1
	var num1 uint8 = math.MaxUint8
	var num2 uint16 = math.MaxUint16
	var num3 uint32 = math.MaxUint32
	var num4 uint64 = math.MaxUint64

	fmt.Println("ex1 :", num1)
	fmt.Println("ex1 :", num2)
	fmt.Println("ex1 :", num3)
	fmt.Println("ex1 :", num4)
	fmt.Println("ex1 :", math.MaxInt8)
	fmt.Println("ex1 :", math.MaxInt16)
	fmt.Println("ex1 :", math.MaxInt32)
	fmt.Println("ex1 :", math.MaxInt64)
	fmt.Println("ex1 :", math.MaxFloat32)
	fmt.Println("ex1 :", math.MaxFloat64)
	fmt.Println("ex1 :", math.MinInt8)
	fmt.Println("ex1 :", math.MinInt16)
	fmt.Println("ex1 :", math.MinInt32)
	fmt.Println("ex1 :", math.MinInt64)

	num5 := 100000 // int
	num6 := int16(10000)
	num7 := uint8(100)

	// fmt.Println("ex2 : ", num5 + num6) // mismatched 예외 발생
	fmt.Println("ex2 : ", num5+int(num6))
	// fmt.Println("ex2 : ", num6 + num7)
	fmt.Println("ex2 : ", num6+int16(num7))
	fmt.Println("ex2 : ", num6 > int16(num7))      // 비교연산도 형이 같아야 비교가능
	fmt.Println("ex2 : ", num6-int16(num7) > 5000) // 산술과 비교 같이

	// 형변환 시 자료 소실
	var n3 int = 12
	var n4 float32 = 8.2
	var n5 uint16 = 1024
	var n6 uint32 = 120000

	// fmt.Println("ex3 : ", n3+n4)
	fmt.Println("ex3 : ", float32(n3)+n4)
	fmt.Println("ex3 : ", n3+int(n4))
	// 주의 20.2이지만 소수점 자리가 소실되면서 20 출력, 실수형 쪽으로 형변환 해야 함
	fmt.Println("ex3 : ", n3+int(n5))
	fmt.Println("ex3 : ", int16(n6)) // 형변환 되면서 값이 이상하게 바뀜

	// 오버플로우 (에러 : 범위 초과)
	// var nn1 uint8 = math.MaxUint8 +1
	// var nn2 uint16 = math.MaxUint16 +1
	// var nn3 uint32 = math.MaxUint32 +1
	// var nn4 uint64 = math.MaxUint64 +1
}
