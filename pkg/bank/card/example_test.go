package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleDeposit_positive() {
	card := IssueCard("TJS", "Black", "Test")
	Deposit(&card, 20_000_00)
	fmt.Println(card.Balance)
	// Output: 2000000
}
func ExampleDeposit_inactive() {
	card := IssueCard("TJS", "Black", "Test")
	card.Active = false
	Deposit(&card, 20_000_00)
	fmt.Println(card.Balance)
	// Output: 0
}
func ExampleDeposit_limit() {
	card := IssueCard("TJS", "Black", "Test")
	Deposit(&card, 50_000_01)
	fmt.Println(card.Balance)
	// Output: 0
}

func ExampleAddBonus_positive() {
	card := IssueCard("TJS", "Black", "Test")
	card.MinBalance = 10_000_00
	card.Balance = 1
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 2466
}
func ExampleAddBonus_inactive() {
	card := IssueCard("TJS", "Black", "Test")
	card.MinBalance = 10_000_00
	card.Balance = 1
	card.Active = false
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 1
}
func ExampleAddBonus_negative() {
	card := IssueCard("TJS", "Black", "Test")
	card.MinBalance = 10_000_00
	card.Balance = -1
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: -1
}
func ExampleAddBonus_limit() {
	card := IssueCard("TJS", "Black", "Test")
	card.MinBalance = 1_000_000_000_00
	card.Balance = 1
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 1
}

func ExamplePaymentSources() {
	cards := []types.Card{
		{
			Active:  true,
			PAN:     "1234",
			Balance: types.Money(100_00),
		},
		{
			Active:  false,
			PAN:     "1234",
			Balance: types.Money(100_00),
		},
		{
			Active:  true,
			PAN:     "5678",
			Balance: types.Money(-100_00),
		},
	}

	sources := PaymentSources(cards)

	for _, source := range sources {
		fmt.Println(source.Number)
	}
	// Output:
	// 1234
}

