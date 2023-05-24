package main

import (
	"fmt"
	"sync"
)

type Income struct {
	Source string
	Value  int
}

var wg sync.WaitGroup

func main() {

	MonthlyIncomes := []Income{
		{Source: "Main Job", Value: 500},
		{Source: "Side Job", Value: 10},
		{Source: "Business", Value: 50},
		{Source: "Other Business", Value: 100},
	}

	var mutex sync.Mutex
	wg.Add(len(MonthlyIncomes))
	var finalBankBalace int

	for i, income := range MonthlyIncomes {

		go func(i int, income Income) {
			/* Defer wg.done() always needs to be inside the go routine call (otherwise how will wg
			know that its completed and you have to decrease the wg counter)*/
			defer wg.Done()
			for week := 0; week < 52; week++ {
				mutex.Lock()
				finalBankBalace += income.Value
				mutex.Unlock()
			}

		}(i, income)

	}

	wg.Wait()

	fmt.Printf("Your yearly income is : Rs.%d.00", finalBankBalace)

}
