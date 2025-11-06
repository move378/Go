// 구조체 익명 선언 시의 구조
package main

import "fmt"

func main() {
	//구조체 익명 선언 및 활용

	//예제1 type 구조체명 타입[struct, func, int 등등] (여기서 익명은 구조체명을 빼버리면 익명)
	car1 := struct{ name, color string }{"520d", "red"}
	// type과 구조체명 생략된 익명 구조체를 바로 선언하고 car1 에 할당!

	fmt.Println("ex1 : ", car1) // car1의 값만 출력
	fmt.Printf("ex1 : %#v\n", car1)
	// %#v #을 쓰면 타입과 속성 그리고 키:밸류 형태로 아래처럼 모든 내용이 나옴(모든 디테일한 명세표가 출력됨)
	// ex1 : struct { name string; color string }{name:"520d", color:"red"}

	//예제2
	cars := []struct{ name, color string }{{"520d", "red"}, {"220d", "white"}, {"420d", "black"}}
	for _, c := range cars { // 위 구조체의 배열(여러 대의 자동차 인스턴스)
		fmt.Printf("(%s, %s) ----- (%#v)\n", c.name, c.color, c)
	}

}
