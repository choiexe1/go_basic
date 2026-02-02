package minibank

import (
	"errors"
	"testing"
)

func setupBank() *Bank {
	bank := NewBank()
	bank.CreateAccount("001", "Jay")
	bank.CreateAccount("002", "Kim")
	return bank
}

func TestDeposit(t *testing.T) {
	bank := setupBank()

	account, _ := bank.FindAccount("001")
	err := account.Deposit(1000, "EXTERNAL")
	if err != nil {
		t.Errorf("에러 = %v, 기대값 = nil", err)
	}

	if account.Balance != 1000 {
		t.Errorf("잔액 = %f, 기대값 = 1000", account.Balance)
	}

	if len(account.Transactions) <= 0 {
		t.Errorf("거래 수 = %v, 기대값 = 1", len(account.Transactions))
	}

	err = account.Deposit(-500, "EXTERNAL")
	if !errors.Is(err, ErrInvalidAmount) {
		t.Errorf("에러 = %v, 기대값 = ErrInvalidAmount", err)
	}
}

func TestWithdraw(t *testing.T) {
	bank := setupBank()
	account, _ := bank.FindAccount("001")
	account.Deposit(5000, "EXTERNAL")

	err := account.Withdraw(3000, "EXTERNAL")
	if err != nil {
		t.Errorf("에러 = %v, 기대값 = nil", err)
	}

	if account.Balance != 2000 {
		t.Errorf("잔액 = %f, 기대값 = 2000", account.Balance)
	}

	err = account.Withdraw(999999, "EXTERNAL")
	if !errors.Is(err, ErrInsufficientBalance) {
		t.Errorf("에러 = %v, 기대값 = ErrInsufficientBalance", err)
	}

	err = account.Withdraw(-100, "EXTERNAL")
	if !errors.Is(err, ErrInvalidAmount) {
		t.Errorf("에러 = %v, 기대값 = ErrInvalidAmount", err)
	}
}

func TestHistory(t *testing.T) {
	bank := setupBank()
	account, _ := bank.FindAccount("001")

	account.Deposit(5000, "EXTERNAL")
	account.Withdraw(2000, "EXTERNAL")

	history := account.History()

	if len(history) != 2 {
		t.Fatalf("거래 수 = %d, 기대값 = 2", len(history))
	}

	if history[0].Type != Deposit {
		t.Errorf("첫 번째 거래 = %s, 기대값 = deposit", history[0].Type)
	}

	if history[1].Type != Withdraw {
		t.Errorf("두 번째 거래 = %s, 기대값 = withdraw", history[1].Type)
	}
}

func TestTransfer(t *testing.T) {
	bank := setupBank()
	jay, _ := bank.FindAccount("001")
	jay.Deposit(10000, "EXTERNAL")

	err := bank.Transfer("001", "002", 3000)
	if err != nil {
		t.Errorf("에러 = %v, 기대값 = nil", err)
	}

	if jay.Balance != 7000 {
		t.Errorf("Jay 잔액 = %f, 기대값 = 7000", jay.Balance)
	}

	kim, _ := bank.FindAccount("002")
	if kim.Balance != 3000 {
		t.Errorf("Kim 잔액 = %f, 기대값 = 3000", kim.Balance)
	}

	err = bank.Transfer("001", "002", 999999)
	if !errors.Is(err, ErrInsufficientBalance) {
		t.Errorf("에러 = %v, 기대값 = ErrInsufficientBalance", err)
	}

	err = bank.Transfer("001", "999", 1000)
	if !errors.Is(err, ErrAccountNotFound) {
		t.Errorf("에러 = %v, 기대값 = ErrAccountNotFound", err)
	}
}
