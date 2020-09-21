package transfer

//Total возвращает итоговую сумму для зачисления перевода
func Total(amount int) (total int) {
	bonusPercent := 5
	total = amount + amount/1000*bonusPercent
	return total
}
