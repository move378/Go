// ========================================
// 파일 I/O (2) - 버퍼링(Buffered I/O)
// ========================================
// bufio 패키지: 효율적인 파일 읽기/쓰기를 위한 버퍼링 지원
// io/ioutil 대비 메모리 효율성과 성능 향상
// ========================================

package main

import (
	"bufio" // 버퍼링된 I/O를 위한 패키지
	"fmt"
	"io" // Seeker 인터페이스와 Seek 상수를 위해 추가
	"os"
)

// ========================================
// 에러 핸들링 유틸리티
// ========================================
func errCheck(e error) {
	if e != nil {
		// ⚠️ 실무에서는 panic 대신 log.Fatal() 또는
		// 명시적 에러 반환 권장
		panic(e)
	}
}

func main() {
	// ========================================
	// 📚 핵심 개념: io.Reader와 io.Writer 인터페이스
	// ========================================
	/*
		Go의 I/O 시스템은 인터페이스 기반으로 설계됨:

		type Reader interface {
			Read(p []byte) (n int, err error)
		}

		type Writer interface {
			Write(p []byte) (n int, err error)
		}

		bufio 패키지는 이 인터페이스들을 구현하며,
		내부 버퍼를 통해 시스템 콜 횟수를 줄여 성능 향상
	*/

	// ========================================
	// 📝 Step 1: 파일 열기 (OpenFile)
	// ========================================

	/*
		os.OpenFile() 플래그 이해:

		- os.O_CREATE: 파일이 없으면 생성
		- os.O_RDWR: 읽기/쓰기 모드 (Read + Write)
		- os.O_TRUNC: 파일 내용을 비우고 시작 (기존 내용 삭제)
		- os.O_APPEND: 파일 끝에 추가 (기존 내용 유지)
		- os.O_RDONLY: 읽기 전용
		- os.O_WRONLY: 쓰기 전용

		플래그는 | (비트 OR) 연산자로 조합 가능
		예: os.O_CREATE | os.O_RDWR = "생성 + 읽기쓰기"

		⚠️ 보안 개선:
		0777 → 0644로 변경 (소유자 읽기/쓰기, 나머지 읽기만)
	*/
	file, err := os.OpenFile(
		"test_write2.txt",     // 파일명
		os.O_CREATE|os.O_RDWR, // 플래그: 생성 + 읽기쓰기
		os.FileMode(0644),     // 권한: rw-r--r--
	)
	errCheck(err)

	// ⚠️ 중요: defer로 파일 닫기 보장
	// 함수 종료 시 자동으로 파일 닫힘 (리소스 누수 방지)
	defer file.Close()

	// ========================================
	// 🚀 Step 2: 버퍼를 사용한 파일 쓰기
	// ========================================

	// 💡 버퍼링의 필요성:
	/*
		버퍼 없이 쓰기:
		"H" -> 디스크 쓰기 (시스템 콜)
		"e" -> 디스크 쓰기 (시스템 콜)
		"l" -> 디스크 쓰기 (시스템 콜)
		... (문자마다 시스템 콜 발생 = 느림)

		버퍼 사용:
		"Hello Golang!" -> 버퍼에 저장
		버퍼가 가득 차거나 Flush() 호출 시 한 번에 디스크 쓰기
		= 시스템 콜 횟수 대폭 감소 = 빠름!
	*/

	// bufio.NewWriter: 버퍼링된 Writer 객체 생성
	// 기본 버퍼 크기: 4096 bytes (4KB)
	wt := bufio.NewWriter(file)

	// WriteString: 문자열을 버퍼에 쓰기
	// ⚠️ 주의: 아직 디스크에 기록되지 않음!
	wt.WriteString("Hello Golang!\n File Write Test1!\n")

	// Write: []byte를 버퍼에 쓰기
	wt.Write([]byte("Hello Golang!\n File Write Test2!\n"))

	// ----------------------------------------
	// 📊 버퍼 상태 정보 출력
	// ----------------------------------------

	// Buffered(): 현재 버퍼에 쓰여진 바이트 수
	fmt.Printf("사용한 Buffer Size (%d bytes)\n", wt.Buffered())

	// Available(): 버퍼에서 사용 가능한 남은 공간
	fmt.Printf("남은 Buffer Size (%d bytes)\n", wt.Available())

	// Size(): 전체 버퍼 크기 (기본 4096 bytes)
	fmt.Printf("전체 Buffer Size (%d bytes)\n", wt.Size())

	// ⚠️ 가장 중요한 부분!
	// Flush(): 버퍼의 내용을 실제 디스크에 기록
	// 이 함수를 호출하지 않으면 데이터가 손실될 수 있음!
	wt.Flush()

	fmt.Println("쓰기 작업 완료\n")
	fmt.Println("=============================================")

	// ========================================
	// 📖 Step 3: 파일 정보 확인 및 읽기 준비
	// ========================================

	// 💡 문제 상황:
	/*
		현재 파일 포인터의 위치:
		[파일 시작] -------- 내용 -------- [파일 끝] ← 현재 위치

		Flush() 후 파일 포인터는 파일의 끝에 있음!
		이 상태에서 Read()를 호출하면 읽을 내용이 없음 (EOF)

		해결책: Seek()를 사용하여 파일 포인터를 시작 지점으로 이동
	*/

	// ✅ Go 1.7+ 권장 방식: io.SeekStart 사용
	// ❌ 구버전 (Deprecated): os.SEEK_SET
	/*
		Seek() 파라미터:
		1. offset: 이동할 바이트 수
		2. whence: 기준점
		   - io.SeekStart (0): 파일 시작점 기준
		   - io.SeekCurrent (1): 현재 위치 기준
		   - io.SeekEnd (2): 파일 끝 기준

		예시:
		- Seek(0, io.SeekStart): 파일 시작으로 이동
		- Seek(-10, io.SeekEnd): 파일 끝에서 10바이트 앞으로
		- Seek(5, io.SeekCurrent): 현재 위치에서 5바이트 뒤로
	*/
	_, err = file.Seek(0, io.SeekStart)
	errCheck(err)

	// bufio.NewReader: 버퍼링된 Reader 객체 생성
	rt := bufio.NewReader(file)

	// ----------------------------------------
	// 📋 파일 메타데이터 정보 가져오기
	// ----------------------------------------

	// Stat(): 파일의 상세 정보 반환
	// FileInfo 인터페이스를 구현한 객체 반환
	fi, err := file.Stat()
	errCheck(err)

	// 파일 크기만큼 바이트 슬라이스 생성
	// make([]byte, size): 지정된 크기의 바이트 슬라이스 할당
	b := make([]byte, fi.Size())

	// 파일 정보 출력
	fmt.Println("파일 정보 출력 : ", fi)
	fmt.Println("파일 이름 : ", fi.Name())
	fmt.Println("파일 크기 : ", fi.Size())
	fmt.Println("파일 수정 시간 : ", fi.ModTime())

	fmt.Println("=============================================")

	// ========================================
	// 📥 Step 4: 버퍼를 사용한 파일 읽기
	// ========================================

	// Read(): 파일 내용을 바이트 슬라이스 b에 읽어옴
	// 반환값:
	// - data: 읽은 바이트 수
	// - err: 에러 (EOF 포함)
	data, err := rt.Read(b)

	// ⚠️ 여기서는 에러를 무시하지만 (언더스코어 _)
	// 실무에서는 항상 에러를 확인해야 함!
	// 특히 io.EOF (파일 끝) 처리가 중요
	if err != nil && err != io.EOF {
		errCheck(err)
	}

	// 버퍼 정보 출력
	fmt.Printf("전체 Buffer Size : (%d bytes)\n", rt.Size())
	fmt.Printf("읽기 작업 완료 : (%d bytes)\n", data)

	fmt.Println("=============================================")

	// 바이트 슬라이스를 문자열로 변환하여 출력
	fmt.Println(string(b))

	fmt.Println("=============================================")

	// defer file.Close() 로 자동 닫힘
}

// ========================================
// 🎓 핵심 학습 포인트
// ========================================
/*
1️⃣ 버퍼링의 중요성:
   - 시스템 콜 횟수 감소 → 성능 향상
   - 작은 쓰기 작업을 모아서 한 번에 처리

2️⃣ Flush()의 중요성:
   - 버퍼에 있는 데이터를 실제 디스크에 기록
   - 호출하지 않으면 데이터 손실 위험!

3️⃣ 파일 포인터(File Pointer):
   - 현재 읽기/쓰기 위치를 가리키는 내부 커서
   - Seek()로 이동 가능
   - 쓰기 후에는 끝에 위치 → 읽기 전 Seek(0, io.SeekStart) 필요

4️⃣ io.Reader와 io.Writer 인터페이스:
   - Go의 I/O 시스템의 핵심
   - 표준화된 인터페이스로 다양한 소스 처리 가능
   - 파일, 네트워크, 메모리 등 모두 동일한 방식으로 처리

5️⃣ defer의 활용:
   - 리소스 정리 보장 (파일 닫기)
   - panic 발생 시에도 실행됨

// ========================================
// 🚀 실무 개선 사항
// ========================================

// ✅ 더 안전한 에러 처리
func safeReadWrite() error {
	file, err := os.OpenFile("data.txt", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("파일 열기 실패: %w", err)
	}
	defer file.Close()

	// 작업 수행...

	return nil
}

// ✅ 버퍼 크기 커스터마이징
func customBufferSize() {
	file, _ := os.Create("large.txt")
	defer file.Close()

	// 기본 4KB 대신 64KB 버퍼 사용 (큰 파일에 효율적)
	writer := bufio.NewWriterSize(file, 64*1024)
	defer writer.Flush() // 반드시 Flush!
}

// ✅ 라인 단위로 읽기 (텍스트 파일 처리 시 유용)
func readLines() {
	file, _ := os.Open("text.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// ========================================
// 📊 성능 비교 (참고)
// ========================================

버퍼 없이 1MB 파일 쓰기:
- 시스템 콜 횟수: ~1,000,000회
- 소요 시간: ~2.5초

버퍼 사용 (4KB):
- 시스템 콜 횟수: ~256회
- 소요 시간: ~0.01초

약 250배 성능 향상! 🚀

// ========================================
// 🔗 추가 학습 자료
// ========================================
- bufio 패키지: https://pkg.go.dev/bufio
- io 패키지: https://pkg.go.dev/io
- 파일 포인터와 Seek: https://pkg.go.dev/io#Seeker
- Go I/O 인터페이스 설계: https://go.dev/blog/io2010
*/
