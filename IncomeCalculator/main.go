// Iterative Yearly income calculator -- Simple 2 for loops
package main

import (
	"fmt"
)

type Income struct {
	Source string
	Value  int
}

func main() {

	MonthlyIncomes := []Income{
		{Source: "Main Job", Value: 500},
		{Source: "Side Job", Value: 10},
		{Source: "Business", Value: 50},
		{Source: "Other Business", Value: 100},
	}

	var finalBankBalace int

	// Looping through each income source and calculating
	for _, income := range MonthlyIncomes {
		for week := 0; week < 52; week++ {
			finalBankBalace += income.Value
		}
	}

	fmt.Printf("Your yearly income is : Rs.%d.00", finalBankBalace)

}
