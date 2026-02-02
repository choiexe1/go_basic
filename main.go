package main

import (
	"errors"
	"fmt"
	minibank "go_basic/cmd/minibank"
)

func main() {
	bank := minibank.NewBank()

	jay, _ := bank.CreateAccount("001", "Jay")
	kim, _ := bank.CreateAccount("002", "Kim")

	jay.Deposit(100000, "EXTERNAL")
	kim.Deposit(50000, "EXTERNAL")
	fmt.Println(jay)
	fmt.Println(kim)

	fmt.Println("--- Transfer ---")
	err := bank.Transfer("001", "002", 30000)
	if err != nil {
		fmt.Println("에러:", err)
	}
	fmt.Println(jay)
	fmt.Println(kim)

	fmt.Println("--- 잔액 부족 ---")
	err = bank.Transfer("001", "002", 999999)
	if err != nil {
		fmt.Println("에러:", err)
	}
	if errors.Is(err, minibank.ErrInsufficientBalance) {
		fmt.Println("원인: 잔액 부족 확인됨")
	}

	fmt.Println("--- 없는 계좌 ---")
	err = bank.Transfer("001", "999", 1000)
	if errors.Is(err, minibank.ErrAccountNotFound) {
		fmt.Println("에러:", err)
	}

	fmt.Println("--- 중복 계좌 ---")
	_, err = bank.CreateAccount("001", "Duplicate")
	if errors.Is(err, minibank.ErrAlreadyExistID) {
		fmt.Println("에러:", err)
	}

	fmt.Println("--- 거래 내역 ---")
	for _, t := range jay.History() {
		fmt.Printf("[%s] %s → %s : %.0f원\n", t.Type, t.From, t.To, t.Amount)
	}
}
