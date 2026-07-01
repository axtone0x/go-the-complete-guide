package main

import (
	"fmt"
)

func main(){
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	// fmt.Print("Revenue: ")
	// fmt.Scan(&revenue)
	revenue := inputAndPrint("Revenue: ")

	// fmt.Print("Expenses: ")
	// fmt.Scan(&expenses)
	expenses := inputAndPrint("Expenses: ")

	// fmt.Print("Tax rate: ")
	// fmt.Scan(&taxRate)
	taxRate := inputAndPrint("Tax Rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n",ebt)
	fmt.Printf("%.1f\n",profit)
	fmt.Printf("%.3f\n",ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt/profit

	return ebt, profit, ratio
}

func inputAndPrint(text string) (value float64) {
	fmt.Print(text)
	fmt.Scan(&value)

	return value
}