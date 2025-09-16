package main

import "fmt"

func main() {
	const (
		_ = iota
		A
		_
		C
	)

	const (
		_ = iota + 0.75*2

		DEFAULT
		SILVER
		_
		PLATINUM
	)

	fmt.Println("D : ", DEFAULT)
	fmt.Println("S : ", SILVER)
	fmt.Println("P : ", PLATINUM)
	fmt.Println("A : ", A, "B : ", " C : ", C)

}
