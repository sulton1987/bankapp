package types

type Money int64
type Currency string
type PAN string

const (
	TJS Currency = "TJS"
	USD Currency = "USD"
	RUB Currency = "RUB"
)

type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
	MinBalance Money
}

type Payment struct {
	ID     int
	Amount Money
}

type PaymentSource struct {
	Type    string
	Number  string
	Balance Money
}
