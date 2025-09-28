// 사용자 정의 타입(3 사용자 정의 함수)
package main

import "fmt"

type totCost func(int, int) int // totCost 타입정의

func describe(cnt int, price int, fn totCost) { //재정의된 totCost를 받지 않으면 호출 자체가 불가능
	fmt.Printf("cnt: %d, price: %d, orderPrice: %d", cnt, price, fn(cnt, price))
}

func main() {
	//함수 사용자 정의 타입

	//예제1
	var orderPrice totCost                      // 선언 -- 정의
	orderPrice = func(cnt int, price int) int { // totCost 형으로 매개변수 형식과 반환 형식이 일치해야 함
		return (cnt * price) + 1000000
	}
	describe(3, 10000000, orderPrice)
}
