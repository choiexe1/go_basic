package main

import (
	"fmt"
	stdinterfaces "go_basic/cmd/16_std_interfaces"
	"io"
)

func main() {
	entry, err := stdinterfaces.NewLogEntry("INFO", "서버 시작됨")
	fmt.Println(entry.String())
	fmt.Println("에러:", err)

	_, err = stdinterfaces.NewLogEntry("DEBUG", "잘못된 레벨")
	fmt.Println("잘못된 레벨 에러:", err)

	writer := &stdinterfaces.LogWriter{}
	fmt.Fprintf(writer, "%s\n", entry.String())

	entry2, _ := stdinterfaces.NewLogEntry("ERROR", "디스크 부족")
	fmt.Fprintf(writer, "%s\n", entry2.String())
	fmt.Println("=== LogWriter 버퍼 내용 ===")
	fmt.Println(writer.String())

	reader := stdinterfaces.NewLogReader([]byte("LogReader로 읽는 테스트 데이터"))
	buf := make([]byte, 10)
	fmt.Println("=== LogReader 읽기 ===")
	for {
		n, err := reader.Read(buf)
		if n > 0 {
			fmt.Print(string(buf[:n]))
		}
		if err == io.EOF {
			break
		}
	}
	fmt.Println()
}
