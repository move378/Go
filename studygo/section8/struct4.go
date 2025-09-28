//구조체 익명 선언시의 구조
package main

import "fmt"

func main() {
	//구조체 익명 선언 및 활용

	//예제1 type 구조체명 타입[struct, func, int 등등]
	car1 := struct{ name, color string }{"520d", "red"}
	// type과 구조체명 생략된 익명 구조체를 바로 선언하고 car1 에 할당!

	fmt.Println("ex1 : ", car1)
	fmt.Printf("ex1 : %#v\n", car1)

	//예제2
	cars := []struct{ name, color string }{{"520d", "red"}, {"220d", "white"}, {"420d", "black"}}
	for _, c := range cars {
		fmt.Printf("(%s, %s) ----- (%#v)\n", c.name, c.color, c)
	}

}
