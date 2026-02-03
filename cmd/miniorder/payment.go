package miniorder

import "sync"

type Payment interface {
	Pay(amount int) error
}

type CreditCard struct {
	CardNumber string
	Balance    int
	mu         sync.Mutex
}

type BankTransfer struct {
	AccountNumber string
	Balance       int
	mu            sync.Mutex
}

func (c *CreditCard) Pay(amount int) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if amount <= 0 {
		return ErrInvalidAmount
	}

	if c.Balance >= amount {
		c.Balance -= amount
		return nil
	}

	return &PaymentError{
		Method: "CreditCard",
		Reason: ErrInsufficientBalance,
	}

}

func (b *BankTransfer) Pay(amount int) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if amount <= 0 {
		return ErrInvalidAmount
	}

	if b.Balance >= amount {
		b.Balance -= amount
		return nil
	}

	return &PaymentError{
		Method: "BankTransfer",
		Reason: ErrInsufficientBalance,
	}
}
