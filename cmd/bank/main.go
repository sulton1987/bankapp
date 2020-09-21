package main

import (
	"bank/pkg/bank/deposit"
	"fmt"
)

func main() {
	min, max := deposit.Calculate(500_000_00, "TJS")
	fmt.Println(min)
	fmt.Println(max)
}
