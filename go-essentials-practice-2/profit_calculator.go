package main

import (
	"errors"
	"fmt"
	"os"
)

func writeCalculationsToFile(ebt, profit, ratio float64) error {
	content := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n",ebt, profit, ratio)
	err := os.WriteFile("calculations.txt", []byte(content),0644)

	if err != nil {
		return errors.New("An error ocurred while trying to write the file")
	}

	return nil
}

func main(){
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	// fmt.Print("Revenue: ")
	// fmt.Scan(&revenue)
	revenue, err := inputAndPrint("Revenue: ")

	if err != nil {
		fmt.Print(err)
		panic("Cannot continue with the execution")
	}

	// fmt.Print("Expenses: ")
	// fmt.Scan(&expenses)
	expenses, err := inputAndPrint("Expenses: ")

	if err != nil {
		fmt.Print(err)
		panic("Cannot continue with the execution")
	}

	// fmt.Print("Tax rate: ")
	// fmt.Scan(&taxRate)
	taxRate, err := inputAndPrint("Tax Rate: ")

	if err != nil {
		fmt.Print(err)
		panic("Cannot continue with the execution")
	}

	// if err1 != nil || err2 != nil || err3 != nil {
	// 	fmt.Print(err1)
	// 	return
	// }

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	if err != nil {
		fmt.Println("ERROR: ", err)
		panic("Can't continue with the execution")
	}

	fmt.Printf("%.1f\n",ebt)
	fmt.Printf("%.1f\n",profit)
	fmt.Printf("%.3f\n",ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt/profit

	err := writeCalculationsToFile(ebt, profit, ratio)

	if err != nil {
		fmt.Println("ERROR: ", err)
		panic("Cannot continue with the execution")
	}

	return ebt, profit, ratio
}

func inputAndPrint(text string) (float64, error) {
	var value float64
	fmt.Print(text)
	fmt.Scan(&value)

	if value <= 0 {
		return 0, errors.New("Value must be positive and greater than 0")
	}

	return value, nil
}