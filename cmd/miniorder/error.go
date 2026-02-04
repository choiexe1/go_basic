package miniorder

import (
	"errors"
	"fmt"
)

var ErrAlreadyRegisteredProduct = errors.New("이미 등록된 상품")
var ErrProductNotFound = errors.New("상품 없음")
var ErrInvalidQuantity = errors.New("수량 ≤ 0")
var ErrOutOfStock = errors.New("재고 부족")
var ErrInvalidAmount = errors.New("금액 ≤ 0")
var ErrInsufficientBalance = errors.New("잔액 부족")

type PaymentError struct {
	Method string
	Reason error
}

func (e *PaymentError) Error() string {
	return fmt.Sprintf("결제 실패 [%s]: %v", e.Method, e.Reason)
}

func (e *PaymentError) Unwrap() error {
	return e.Reason
}
