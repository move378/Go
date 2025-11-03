// 데이터타입 : Numeric (1)
package main

import "fmt"

func main() {
	/*
		데이터 타입 : 숫자형
		정수, 실수, 복소수
		32bit, 64bit, unsigned(양수)
		정수 : 8진수(0), 16진수(0x), 10진수
		정확하게 명시적으로 선언해줘야 합니다.
	*/

	var num1 int = 17
	var num2 int = -68
	var num3 int = 0631
	var num4 int = 0x32fa2c75

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)

	/*
		uint8 = Byte, int32 = rune, uint = either 32 or 64 bits, int = same size as uint
		uintptr = an unsigned integer large enough to store the uninterpreted bits of a pointer value
		uintptr = 포인터 값의 해석되지 않은 비트들을 저장하기에 충분히 큰 부호 없는 정수
	*/

	//예제2 - 아스키(영문)
	var char1 byte = 72 // byte==uint8, rune == int32
	var char2 byte = 0110
	var char3 byte = 0x48

	//유니코드(한글)
	var char4 rune = 50556   // 유니코드
	var char5 rune = 0142574 // 44032(8진수)
	var char6 rune = 0xC57C  // 44032 (16진수)

	fmt.Printf("%c %c %c\n ", char1, char2, char3) // %char, %dex, %oxigen, %x(hex) 16진수
	fmt.Printf("%d %d %d\n ", char1, char2, char3)
	fmt.Printf("%c %o %x\n ", char1, char2, char3)

	fmt.Printf("%c %c %c\n ", char4, char5, char6)
	fmt.Printf("%d %d %d\n ", char4, char5, char6)
	fmt.Printf("%d %o %x\n ", char4, char5, char6)
	// D dex C char

	//예제3 - 부동소수점
	//float32(7자리), float64(15자리)

	var float1 float32 = 0.14
	var float2 float32 = .75647
	var float3 float32 = 442.0378373
	var float4 float32 = 10.0

	//지수 표기법
	var float5 float32 = 14e6
	var float6 float64 = .156875e+3
	var float7 float64 = 5.32521e-10

	fmt.Println("num1 : ", float1)
	fmt.Println("num2 : ", float2)
	fmt.Println("num3 : ", float3)
	fmt.Println("num4 : ", float4)
	fmt.Println("num4-0.1 nomal : ", float4-0.1)
	fmt.Println("num4-0.1 float32: ", float32(float4-0.1))
	fmt.Println("num4-0.1 float64 : ", float64(float4-0.1))
	// 형 변환 시 부동소수점 오류! 엄연히 다른 숫자 돈이라고 한다면?
	fmt.Println("ex1 : ", float5)
	fmt.Println("ex1 : ", float6)
	fmt.Println("ex1 : ", float7)
}
