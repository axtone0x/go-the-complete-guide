package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5
// var investmentAmount float64
// var years float64 = 10.0
// var expectedReturnRate float64 = 5.5

func main() {
	var investmentAmount float64
	years := 10.0
	expectedReturnRate := 5.5

	//Using the funcion
	outputText("Investment amount: ")
	fmt.Scan(&investmentAmount)

	outputText("years: ")
	fmt.Scan(&years)

	outputText("expectedReturnRate: ")
	fmt.Scan(&expectedReturnRate)

	//Calculation
	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100,years)

	//Calculations using a function
	fv, frv := calculateFutureValue(investmentAmount,expectedReturnRate,years)

	//Storing a string in a variable and then print it
	// formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	// formattedRFV := fmt.Sprintf("Future Real Value: %.1f\n",futureRealValue)
	// fmt.Print(formattedFV,formattedRFV)

	//Printing info formatted
	//fmt.Printf("Future value: %.1f\nFuture Real Value: %.1f",futureValue, futureRealValue)
	//fmt.Println(futureRealValue)
	fmt.Printf("Future value: %.1f\nFuture Real Value: %.1f",fv, frv)

	// fmt.Printf(`
	// Future Value: %.1f
	// Future Real Value: %.1f`,futureValue,futureRealValue)
}

//functions
//--simple function--
func outputText(text string) {
	fmt.Print(text)
}

//----
// func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
// 	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
// 	futureRealValue := futureValue / math.Pow(1+inflationRate/100,years)

// 	return futureValue, futureRealValue
// }

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, frv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	frv = fv / math.Pow(1+inflationRate/100,years)
	//You can return values like this
	//return fv, frv -> better and clearer approach
	//Alternative
	return
}



