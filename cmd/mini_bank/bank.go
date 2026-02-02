package minibank

import "fmt"

type Bank struct {
	accounts map[string]*Account
}

func NewBank() *Bank {
	return &Bank{
		accounts: map[string]*Account{},
	}
}

func (b *Bank) CreateAccount(id string, name string) (*Account, error) {
	_, err := b.FindAccount(id)
	if err == nil {
		return nil, ErrAlreadyExistID
	}

	account := &Account{
		ID:           id,
		Name:         name,
		Transactions: []Transaction{},
	}

	b.accounts[id] = account

	return account, nil
}

func (b *Bank) FindAccount(id string) (*Account, error) {
	account, ok := b.accounts[id]

	if ok {
		return account, nil
	}

	return nil, ErrAccountNotFound
}

func (b *Bank) Transfer(from, to string, amount float64) error {
	fromAccount, err := b.FindAccount(from)
	if err != nil {
		return ErrAccountNotFound
	}

	toAccount, err := b.FindAccount(to)
	if err != nil {
		return ErrAccountNotFound
	}

	if err := fromAccount.Withdraw(amount, to); err != nil {
		return fmt.Errorf("Transfer 실패: %w", err)
	}

	toAccount.Deposit(amount, from)

	return nil
}
