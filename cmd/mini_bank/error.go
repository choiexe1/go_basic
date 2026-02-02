package minibank

import "errors"

var ErrInsufficientBalance = errors.New("잔액 부족")
var ErrInvalidAmount = errors.New("금액은 0보다 커야 합니다")
var ErrAccountNotFound = errors.New("계좌를 찾을 수 없습니다")
var ErrAlreadyExistID = errors.New("이미 존재하는 계좌 ID입니다")
