package projects

import (
	"fmt"
	"sync"
)

type Income struct {
	Source string
	Amount int
}

func IncomeCalculator() {
	var (
		bankBalance int
		wg          sync.WaitGroup
		mx          sync.Mutex
	)

	fmt.Printf("initial bank balance: %d\n", bankBalance)

	incomes := []Income{
		{Source: "Main job", Amount: 500},
		{Source: "Part-time job", Amount: 150},
		{Source: "Investments", Amount: 15},
		{Source: "Gifts", Amount: 20},
	}

	wg.Add(len(incomes))

	for i, income := range incomes {

		go func(i int, income Income) {
			defer wg.Done()

			for week := 1; week <= 52; week++ {
				mx.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				mx.Unlock()
				fmt.Printf("on week %d you earned $%d.00 from %s\n", week, income.Amount, income.Source)

			}

		}(i, income)

	}

	wg.Wait()

	fmt.Printf("Balance is: $%d.00\n", bankBalance)

}
