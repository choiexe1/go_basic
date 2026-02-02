package minibank

import "fmt"

type Account struct {
	ID           string
	Name         string
	Balance      float64
	Transactions []Transaction
}

func (a *Account) Deposit(amount float64, from string) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	a.Balance += amount

	a.Transactions = append(a.Transactions, Transaction{
		Type:   Deposit,
		From:   from,
		To:     a.ID,
		Amount: amount,
	})

	return nil
}

func (a *Account) Withdraw(amount float64, to string) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	if amount > a.Balance {
		return ErrInsufficientBalance
	}

	a.Balance -= amount
	a.Transactions = append(a.Transactions, Transaction{
		Type:   Withdraw,
		From:   a.ID,
		To:     to,
		Amount: amount,
	})

	return nil
}

func (a *Account) History() []Transaction {
	return a.Transactions
}

func (a *Account) String() string {
	return fmt.Sprintf("%s (ID: %s) ㅡ 잔액: %f", a.Name, a.ID, a.Balance)
}
