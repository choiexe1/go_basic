package customerror

import (
	"errors"
	"fmt"
	"strings"
)

var EmptyInputError = errors.New("Empty Input")
var InvalidFormatError = errors.New("Invalid Format")

type ParseError struct {
	Line    int
	Content string
	Err     error
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("line %d: %s - %s", e.Line, e.Content, e.Err)
}

func (e *ParseError) Unwrap() error {
	return e.Err
}

func Parse(input string) (map[string]string, error) {
	v := strings.TrimSpace(input)
	m := map[string]string{}

	if v == "" {
		return nil, EmptyInputError
	}

	lines := strings.Split(v, "\n")

	for i, line := range lines {
		parts := strings.SplitN(line, "=", 2)

		if len(parts) != 2 {
			return nil, &ParseError{
				Line:    i + 1,
				Content: line,
				Err:     InvalidFormatError,
			}
		}

		m[parts[0]] = parts[1]
	}

	return m, nil
}
