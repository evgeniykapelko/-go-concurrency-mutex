package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	// variable for bank balance
	var bankBalance int
	var balance sync.Mutex

	// print out starting values
	fmt.Printf("Initial bank Balance: $%d.00\n", bankBalance)
	fmt.Println()

	// define weekly revenue
	income := []Income{
		{Source: "A", Amount: 100},
		{Source: "B", Amount: 200},
		{Source: "C", Amount: 300},
		{Source: "D", Amount: 400},
	}

	// loop through 52 weeks and print out how much is made;
	// keep a running total
	for index, incomeItem := range income {
		wg.Add(1)
		go func(index int, incomeItem Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += incomeItem.Amount
				bankBalance = temp
				balance.Unlock()

				fmt.Printf(
					"On week %d, you earned $%d.00 from %s\n",
					week,
					incomeItem.Amount,
					incomeItem.Source,
				)
			}
		}(index, incomeItem)
	}

	wg.Wait()
	// print out final balance
	fmt.Printf("Final Bank Balance: $%d.00\n", bankBalance)
	fmt.Println()
}
