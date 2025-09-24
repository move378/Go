//구조체 기본(3)
package main

import "fmt"

type car struct {
	color string
	name  string
}

func main() {

	//예제1
	//함수의 매개변수로 전달시에는 기본적으로 값 복사 이므로, 필요시 포인터로 전달해야 한다.
	c1 := car{"red", "220d"}
	c2 := new(car)
	c2.color, c2.name = "white", "sonata"
	c3 := &car{}
	c4 := &car{"black", "520d"}

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c4)
	fmt.Println()
}
