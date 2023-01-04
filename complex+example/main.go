package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type income struct {
	Source string
	Amount int
}

func main() {
	// variable for bank balance
	var bankBalance int

	// print out starting values
	fmt.Printf("initial accoutn balance: $%d.00", bankBalance)
	fmt.Println()

	// define weekly revenue
	incomes := []income{
		{Source: "main job", Amount: 500},
		{Source: "gifts", Amount: 10},
		{Source: "part time job", Amount: 50},
	}
	// loop through 52 weeks and print out how much is made; keep a running total

	// print out final balance
}
