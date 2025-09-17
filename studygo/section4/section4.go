// 패키지(1)
package main

//선언 방법1
/*
import "fmt"
import "os"
*/

//선언 방법2
import (
	"fmt"
	"os"
	"section4/lib"          //빈 식별자 사용
	testlib "section4/lib2" //별칭 사용
)

func main() {

	/*
		패키지 종류
		1.메인 프로그램(main)
		2.다른 패키지에서 호출 가능한 라이브러리

		패키지 : 코드 구조화 및 재사용
		응집도, 결합도
		Go : 패키지 단위의 독립적이고 작은 단위로 개발
		-> 작은 패키지를 결합해서 프로그램을 작성 할 것
		서로 다른 패키지간에 서로 import 후 사용
		패키지 이름 = 디렉터리 이름
		같은 패키지 내 -> 소스파일들은 디렉토리명을 패키지 명으로 사용한다.
		네이밍 규칙 : 소문자 private (scope : 패키지 내부 - 지역), 대문자 Public (전역)
		go : main 패키지는 특별하게 인식
		-> 컴파일러가 공유 라이브러리가 아닌 프로그램의 시작점 start entry point 로 인식
	*/

	var name string

	fmt.Print("이름은? :  ")

	fmt.Scanf("%s", &name)

	fmt.Fprintf(os.Stdout, "Hi %s\n", name)

	fmt.Print("10 보다 큰 수? :  ", lib.CheckNum(15))

	//패키지 접근제어
	//변수,상수,함수,메서드,구조체 등 식별자
	//대문자 : 패키지 외부 접근 가능
	//소문자 : 패키지 외부 접근 불가(패키지 내에서만 접근 가능)

	fmt.Print("100 보다 큰 수? :  ", lib.CheckNum2(101))
	fmt.Print("1000 보다 큰 수? :  ", lib.checkNum3(1001)) //예외 발생 (접근 불가)

	//패키지 접근제어
	//별칭 사용
	//빈 식별자 사용

	fmt.Print("100 보다 큰 수? :  ", testlib.CheckNum1(101))

}
