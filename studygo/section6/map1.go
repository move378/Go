// 자료형 : 맵(1)
package main

import "fmt"

func main() {
	/* 맵(Map)
		맵 : 해시테이블, 딕셔너리(파이썬), Key-Value로 자료 저장
		1) 레퍼런스 타입!(참조 값 전달)
		2) Key-Value되어 있고,
		3) comparable 타입을 키(Key)로 사용 가능, 값(Value)은 모든 타입을 사용 가능하다.
		[참조 타입 중에서도 non-comparable (slice, map, function)이 키로 사용 불가능하다]
		비교 연산자 사용 불가능(*"비교의 의미가 명확하지 않기 때문"**입니다!)
		make 함수 및 축약(리터럴)로 초기화 가능
		순서 없음, 슬라이스나 배열은 순서가 있지만, 맵 데이터 구조는 순서가 없는 구조

	## 🧩 Comparable의 본질: "동등성 정의 가능 여부"

	### **핵심 질문:**
	> "두 값이 '같다'는 것을 명확하게 정의할 수 있는가?"
	```
	Comparable의 조건:
	├─ 1. 비교 연산(==)이 의미가 있어야 함
	├─ 2. 비교 결과가 예측 가능해야 함
	├─ 3. 해시값 계산이 가능해야 함 (Map Key용)
	└─ 4. 불변하거나, 변해도 비교 규칙이 일관돼야 함

	3차원 프레임워크 확장:

	Comparable 타입:
	├─ 불변 참조: string (안전하게 내용 비교)
	├─ 값 타입: int, array (명확한 값 비교)
	└─ 주소 참조: pointer (간단한 주소 비교)

	Non-Comparable 타입:
	├─ 가변 참조 + 복잡: slice, map
	│   └─ 문제: 비교 의미 모호, 비용 큼
	└─ 비교 무의미: function
	    └─ 문제: 같다는 게 무슨 의미?
	```

	### **Map Key 규칙:**
	```
	Map Key 되려면:
	✅ Comparable (== 가능)
	✅ Hashable (hash() 가능)
	✅ 안정적 (불변 또는 변해도 일관성)

	→ Non-Comparable은 조건 1 실패 → Key 불가!
	*/

	//예제1
	var map1 map[string]int = make(map[string]int) //정석 map[KeyType]ValueType
	var map2 = make(map[string]int)                //자료형 생략
	map3 := make(map[string]int)                   //리터럴(축약) 형

	fmt.Println("ex1 : ", map1)
	fmt.Println("ex1 : ", map2)
	fmt.Println("ex1 : ", map3)
	fmt.Println()

	//예제2
	map4 := map[string]int{} // Json 형태로 {}에 묶여서 데이터가 들어감
	map4["apple"] = 25
	map4["banana"] = 40
	map4["orange"] = 33

	map5 := map[string]int{ // 보통 이런 형태를 사용함 보기가 좀 더 편함
		"apple":  15,
		"banana": 30,
		"orange": 23, //콤마 주의
	}

	map6 := make(map[string]int, 10)
	map6["apple"] = 5
	map6["banana"] = 10
	map6["orange"] = 15

	fmt.Println("ex2 : ", map4)
	fmt.Println("ex2 : ", map5)
	fmt.Println("ex2 : ", map6)
	fmt.Println("ex2 : ", map6["orange"])
	fmt.Println("ex2 : ", map6["apple"])
	fmt.Println()

	//맵 조회 및 순회(Iterator)

	//예제3
	map7 := map[string]string{
		"daum":   "http://daum.net",
		"naver":  "http://naver.com",
		"google": "http://google.com", // Json Map 형식은 마지막에 컴마(,)를 찍어줘야 합니다.
	}
	fmt.Println("ex1 : ", map7["google"])
	fmt.Println("ex1 : ", map7["daum"])
	fmt.Println()

	//예제4(순서 없으므로 랜덤)
	for k, v := range map7 {
		fmt.Println("ex2 : ", k, v)
	}

	fmt.Println()

	for _, v := range map7 { //변수 선언 후 사용안하면 예외 발생
		fmt.Println("ex2 : ", v)
	}

	//맵 값 변경 및 삭제

	//예제5(수정)
	map8 := map[string]string{
		"daum":   "http://daum.net",
		"naver":  "http://naver.com",
		"google": "http://google.com",
		"home1":  "http://test1.com", // 마지막에 콤마를 찍어줘야 마무리가 됩니다.
	}
	fmt.Println("ex1 : ", map8)
	map8["home2"] = "http://test2.com" //추가
	fmt.Println("ex1 : ", map8)
	map8["home2"] = "http://test2-2.com" //같은 키가 존재한다면 수정(그냥 대입)
	fmt.Println("ex1 : ", map8)
	fmt.Println()

	//예제6(삭제)
	delete(map8, "home2") // 대상과 키 값을 제공하면 삭제
	fmt.Println("ex1 : ", map8)
	delete(map8, "home1")
	fmt.Println("ex1 : ", map8)

	//맵(Map)
	//맵 조회 할 경우 주의 할점

	//예제7
	map9 := map[string]int{ //int : 0, string : "", float : 0.0
		"apple":  15,
		"banana": 115,
		"orange": 1115,
		"lemon":  0,
	}

	value1 := map9["lemon"]
	value2 := map9["kiwi"]     // 해당 map 타입(int)의 초기화 값, 0을 반환!
	value3, ok := map9["kiwi"] // map의 경우 기본적으로 두개의 값을 리턴합니다. 해당 값과 키의 존재 유무!
	// go의 맵에서는 변수를 두개를 줘서 해당 키의 존재 유무를 확인하는 것이 좋습니다.

	fmt.Println("ex1 : ", value1)
	fmt.Println("ex1 : ", value2)
	fmt.Println("ex1 : ", value3, ok) //두 번째 리턴 값으로 키 존재 유무 확인
	fmt.Println()

	//예제2
	if value, ok := map9["kiwi"]; ok {
		fmt.Println("ex2 : ", value)
	} else {
		fmt.Println("ex2 : kiwi is not exist!")
	}

	if value, ok := map9["banana"]; ok {
		fmt.Println("ex2 : ", value)
	} else {
		fmt.Println("ex2 : banana is not exist!")
	}

	if _, ok := map9["kiwi"]; !ok { // value를 받지 않고, 키가 있는지 없는지만 확인 (else문을 안쓰려고)
		fmt.Println("ex2 : kiwi is not exist!")
	}

	//다른 언어와 다른 부분 때문에 헤멜 수가 있으니 이런 세부적인 차이를 알아야,
	//어디가 잘못되었는지, 왜 안되는지 파악이 가능하다.
}
