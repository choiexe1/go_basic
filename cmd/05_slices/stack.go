package stack

import "errors"

var stack []int = make([]int, 0)

func Push(value int) {
	stack = append(stack, value)
}

func Pop() (int, error) {
	if len(stack) == 0 {
		return -1, errors.New("stack is empty")
	}

	value := stack[len(stack)-1]
	stack = stack[:len(stack)-1]

	return value, nil
}

func Peek() (int, error) {
	if len(stack) == 0 {
		return -1, errors.New("stack is empty")
	}

	value := stack[len(stack)-1]

	return value, nil
}

func Size() int {
	return len(stack)
}
