// 데이터 타입 : 문자열(1)

package main

import (
	"fmt"
	"unicode/utf8"
	_ "unicode/utf8"
)

func main() {
	/*
		큰 따옴표 "", 백스쿼트 ``
		Go Lang : 문자 char '' 타입 존재하지 않음 -> rune(int32)로 문자 코드 값으로 표현
		문자 : '' 로 작성
		자주 사용하는 escape : \\ \' \", \a(콘솔벨), \b(백스페이스), \f(쪽바꿈), \n(줄바꿈), \r(복귀), \t(탭), \u(유니코드)
	*/

	//예제1 escape로 문자 치환
	var str1 string = "c:\\go_study\\src\\"
	fmt.Println(str1)
	str2 := `c:\go_study\src\` // 백틱 기호를 스면 문자 그대로
	fmt.Println(str2)

	//예제2
	var str3 string = "Hello, world!"
	var str4 string = "안녕하세요." //len, length
	var str5 string = "\ud55c\uae00"

	fmt.Println()
	fmt.Println("ex2 : ", str3)
	fmt.Println("ex2 : ", str4)
	fmt.Println("ex2 : ", str5)

	//예제3 길이(바이트 수)
	fmt.Println("ex3 : ", len(str3))
	fmt.Println("ex3 : ", len(str4)) //한글 3바이트 + 1 byte = 16 byte

	//예제4 길이(실제 문자수)
	fmt.Println("ex4 : ", utf8.RuneCountInString(str3))
	fmt.Println("ex4 : ", utf8.RuneCountInString(str4))
	fmt.Println("ex4 : ", len([]rune(str4)))
	//rune 형식으로 3바이트씩 끊어서 배열이 6개, 문자는 근본이 string 배열임

	/*
		문자열 : 기본 UTF-8 인코딩 (유니코드 문자 집합) -> 바이트 수 주의!
	*/

	// 예제5
	var str6 string = "Golang"
	var str7 string = "World"
	var str8 string = "고프로그래밍"

	fmt.Println("ex1 : ", str6[0], str6[1], str6[2], str6[3], str6[4], str6[5])
	fmt.Println("ex1 : ", str7[0], str7[1], str7[2], str7[3], str7[4])
	fmt.Println("ex1 : ", str8[0], str8[1], str8[2], str8[3], str8[4], str8[5])
	fmt.Println("ex1 : ", str8[0], str8[1], str8[2], str8[3], str7[4], str8[5], str8[len(str1)-2])

	//예제2
	fmt.Printf("ex2 : %c %c %c %c %c %c\n", str6[0], str6[1], str6[2], str6[3], str6[4], str6[5])
	fmt.Printf("ex2 : %c %c %c %c %c\n", str7[0], str7[1], str7[2], str7[3], str7[4])
	fmt.Printf("ex2 : %c %c %c %c %c %c\n", str8[0], str8[1], str8[2], str8[3], str8[4], str8[5])
	//한글 깨짐 발생 시 아래의 방법으로

	conStr := []rune(str8)
	fmt.Printf("ex2 : %c %c %c %c %c %c\n", conStr[0], conStr[1], conStr[2], conStr[3], conStr[4], conStr[5])

	//예제3
	for i, char := range str6 {
		fmt.Printf("ex3 : %c(%d)\t", char, i)
	}

	fmt.Println()

	for i, char := range str7 {
		fmt.Printf("ex3 : %c(%d)\t", char, i)
	}
}
