package enum

type TransactionType int

const (
	topup TransactionType = iota + 1
	withdraw
	transfer
)
