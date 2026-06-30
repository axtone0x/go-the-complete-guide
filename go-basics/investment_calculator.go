package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	years := 10.0
	expectedReturnRate := 5.5

	fmt.Print("Investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("years: ")
	fmt.Scan(&years)

	fmt.Print("expectedReturnRate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100,years)

	// formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	// formattedRFV := fmt.Sprintf("Future Real Value: %.1f\n",futureRealValue)

	// //fmt.Printf("Future value: %.1f\nFuture Real Value: %.1f",futureValue, futureRealValue)
	// //fmt.Println(futureRealValue)

	// fmt.Print(formattedFV,formattedRFV)
}

