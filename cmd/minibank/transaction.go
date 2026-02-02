package minibank

type TransactionType string

const (
	Deposit  TransactionType = "deposit"
	Withdraw TransactionType = "withdraw"
)

type Transaction struct {
	Type   TransactionType
	From   string
	To     string
	Amount float64
}
