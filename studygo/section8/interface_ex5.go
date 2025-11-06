// 인터페이스 고급(5 사용자 정의 구조체(타입)에 대한 타입검사: 포인터형 여부 판별해서 분기하기)
package main

import (
	"fmt"
	"strconv"
)

type Account struct {
	number   string
	balance  float64
	interest float64
}

func structToMsg(arg interface{}) string {
	// 사용자 정의 타입(구조체)가 일반 구조체인지 포인터형 구조체인지 아닌지를 판별해서 분기하는 코드
	switch arg.(type) {
	case Account:
		o := arg.(Account)
		return "Account : " + o.number + " : " + strconv.FormatFloat((o.balance*o.interest+o.balance), 'f', -1, 64)
	case *Account:
		o := arg.(*Account)
		return "*Account : " + o.number + " : " + strconv.FormatFloat((o.balance*o.interest+o.balance), 'f', -1, 64)
	default:
		return "Error"
	}
}

func main() {

	//예제1
	fmt.Println(structToMsg(Account{number: "245-901", balance: 10000000, interest: 0.015}))
	fmt.Println(structToMsg(&Account{number: "245-902", balance: 12000000, interest: 0.035}))

	var user = new(Account)
	user.number = "245-903"
	user.balance = 15000000
	user.interest = 0.025

	fmt.Println(structToMsg(user))
}
