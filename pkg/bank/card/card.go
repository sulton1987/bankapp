package card

import (
	"bank/pkg/bank/types"
)

//AddBonus даёт бонус на остаток в месяц
func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {
	if !card.Active {
		return
	}
	if card.Balance <= 0 {
		return
	}
	bonus := types.Money(int(card.MinBalance) * percent * daysInMonth / 100 / daysInYear)
	if bonus > 5_000_00 {
		return
	}
	card.Balance += bonus
}

//Deposit пополняет баланс карты
func Deposit(card *types.Card, amount types.Money) {
	// depositLimit := types.Money(50_000_00)
	if !card.Active {
		return
	}
	if amount <= 0 {
		return
	}
	if amount > 50_000_00 {
		return
	}
	card.Balance += amount
}

//Withdraw снимает деньги с баланса карты
func Withdraw(card types.Card, amount types.Money) types.Card {
	if !card.Active {
		return card
	}
	if card.Balance < amount {
		return card
	}
	if amount < 0 {
		return card
	}
	if amount > 20_000_00 {
		return card
	}
	card.Balance -= amount

	return card
}

//IssueCard создаёт карту с заданными параметрами
func IssueCard(currency types.Currency, color string, name string) types.Card {
	card := types.Card{
		ID:         1000,
		PAN:        "5058 xxxx xxxx 0001",
		Balance:    0,
		Currency:   currency,
		Color:      color,
		Name:       name,
		Active:     true,
		MinBalance: 0,
	}
	return card
}

//Total возвращает сумму денег активных карт с положиетльным балансом из слайса карт
func Total(cards []types.Card) types.Money {
	total := types.Money(0)
	for _, card := range cards {
		if card.Active && card.Balance > 0{
			total += card.Balance
		}
	}
	return total
}

