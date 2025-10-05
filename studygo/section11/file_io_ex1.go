/* Go 1.16부터 io/ioutil 패키지가 deprecated

ioutil.ReadFile -> os.ReadFile
ioutil.WriteFile -> os.WriteFile
ioutil.ReadAll -> io.ReadAll
ioutil.ReadDir -> os.ReadDir

현재 코드에서는:
ioutil.WriteFile -> os.WriteFile
ioutil.ReadFile -> os.ReadFile

========================================
파일 I/O (입출력) - 최신 버전
========================================
Go 1.16 이전: io/ioutil 패키지 사용 (현재 Deprecated)
Go 1.16 이후: os, io 패키지로 기능 이관
========================================

*/

package main

// os (ioutil) 사용 한줄로! 원자적 연산을 지원함!

import (
	"fmt"
	"os" // 파일 시스템 관련 작업을 위한 표준 라이브러리
)

// ========================================
// errCheck: 에러 핸들링을 위한 유틸리티 함수
// ========================================
// 💡 Go의 관용적 에러 처리 패턴
// - Go는 예외(exception) 대신 명시적 에러 반환을 사용
// - panic은 복구 불가능한 치명적 에러에만 사용 권장
func errCheck(e error) {
	if e != nil {
		// panic: 프로그램 실행을 즉시 중단하고 에러 출력
		// 실무에서는 log.Fatal() 또는 구조화된 에러 처리 권장
		panic(e)
	}
}

func main() {
	// ========================================
	// 1. 파일 쓰기 (Write)
	// ========================================

	// 📝 쓰기할 데이터 준비
	// \n: 줄바꿈 (개행) 문자
	s := "Hello Golang!\n File Write Test!\n"

	// ----------------------------------------
	// 🔐 파일 권한 (Permission) 이해하기
	// ----------------------------------------
	// Unix/Linux 파일 권한 시스템:
	// - 읽기(Read): 4
	// - 쓰기(Write): 2
	// - 실행(Execute): 1
	//
	// 세 자리 숫자의 의미:
	// - 첫째 자리: 소유자(Owner) 권한
	// - 둘째 자리: 그룹(Group) 권한
	// - 셋째 자리: 기타 사용자(Others) 권한
	//
	// 예시: 0644
	// - 0: 8진수 표기 (Go에서 파일 권한은 8진수 사용)
	// - 6 (소유자): 4(읽기) + 2(쓰기) = 읽기/쓰기 가능
	// - 4 (그룹): 읽기만 가능
	// - 4 (기타): 읽기만 가능
	//
	// 일반적인 권한값:
	// - 0644: 일반 파일 (소유자 읽기/쓰기, 나머지 읽기만)
	// - 0755: 실행파일 (소유자 읽기/쓰기/실행, 나머지 읽기/실행)
	// - 0600: 민감한 파일 (소유자만 읽기/쓰기)
	// ----------------------------------------

	// ✅ 최신 방식: os.WriteFile 사용
	// func WriteFile(name string, data []byte, perm FileMode) error
	//
	// 매개변수:
	// 1. name: 생성할 파일 경로/이름
	// 2. data: 쓸 데이터 ([]byte 타입 필요)
	// 3. perm: 파일 권한 (FileMode)
	//
	// 동작 방식:
	// - 파일이 없으면 생성
	// - 파일이 있으면 덮어쓰기 (기존 내용 삭제)
	// - 원자적(atomic) 연산: 쓰기가 완전히 성공하거나 완전히 실패
	err := os.WriteFile("test_write1.txt", []byte(s), os.FileMode(0644))
	errCheck(err)

	// ========================================
	// 2. 파일 읽기 (Read)
	// ========================================

	// ✅ 최신 방식: os.ReadFile 사용
	// func ReadFile(name string) ([]byte, error)
	//
	// 반환값:
	// 1. []byte: 파일의 전체 내용 (바이트 슬라이스)
	// 2. error: 에러 객체 (성공시 nil)
	//
	// 특징:
	// - 파일 전체를 메모리에 로드
	// - 작은 파일(<몇 MB)에 적합
	// - 큰 파일은 bufio.Reader 사용 권장
	// - 파일 핸들 자동 관리 (열기/닫기)
	data, err := os.ReadFile("sample.txt")
	errCheck(err)

	// ========================================
	// 3. 결과 출력
	// ========================================

	fmt.Println("=====================================")

	// []byte를 문자열로 변환하여 출력
	// string(data): 바이트 슬라이스를 UTF-8 문자열로 디코딩
	fmt.Println(string(data))

	fmt.Println("=====================================")
}

// ========================================
// 📚 추가 학습 포인트
// ========================================
//
// 1️⃣ 대용량 파일 처리:
//    - os.Open() + bufio.Reader 사용
//    - 스트리밍 방식으로 메모리 효율적 처리
//
// 2️⃣ 파일 존재 여부 확인:
//    - os.Stat(filename)
//    - errors.Is(err, os.ErrNotExist)
//
// 3️⃣ 안전한 파일 쓰기:
//    - 임시 파일 생성 -> 쓰기 -> 원본 파일과 교체
//    - os.CreateTemp() + os.Rename() 활용
//
// 4️⃣ 파일 추가(Append) 모드:
//    - os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
//
// 5️⃣ 에러 처리 개선:
//    - errors.Is(), errors.As() 활용
//    - 구조화된 에러 반환 (함수가 error 반환)
//
// 🔗 공식 문서:
// - os 패키지: https://pkg.go.dev/os
// - io 패키지: https://pkg.go.dev/io
// - 파일 권한: https://pkg.go.dev/io/fs#FileMode
// ========================================
