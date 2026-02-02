package main

import (
	"errors"
	"fmt"
	customerror "go_basic/cmd/11_errors"
)

func main() {
	result, err := customerror.Parse("host=localhost\nport=8080\ndb=postgres")
	if err != nil {
		fmt.Println("에러:", err)
	} else {
		fmt.Println("정상:", result)
	}

	_, err = customerror.Parse("")
	if errors.Is(err, customerror.EmptyInputError) {
		fmt.Println("빈 입력 감지:", err)
	}

	_, err = customerror.Parse("host=localhost\nbad_line\nport=8080")
	if errors.Is(err, customerror.InvalidFormatError) {
		fmt.Println("포맷 에러:", err)
	}

	var parseErr *customerror.ParseError
	if errors.As(err, &parseErr) {
		fmt.Printf("상세 — %d번째 줄: '%s'\n", parseErr.Line, parseErr.Content)
	}
}
