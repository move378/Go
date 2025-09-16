// 변수2
package main

import "fmt"

func main() {
	// 여러 개 선언 [자주 사용]
	var (
		name      string = "machine"
		height    int32
		weight    float32
		isRunning bool
	)

	height = 250
	weight = 350.56
	isRunning = true

	fmt.Println("name : ", name, "height : ", height, "weight : ", weight, "isRunning", isRunning)

}
