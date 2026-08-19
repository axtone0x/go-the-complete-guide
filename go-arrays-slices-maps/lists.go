package main

import "fmt"

type Product struct {
	title string
	id string
	price float64
}

type TemperatureData struct {
	day1 float64
	day2 float64
}

func main() {
	//var productNames [4]string = [4]string{"Spoon"}
	prices := [7]float64{100.1,200.2,300.3,400.4,500.5,600.6,700.7}
	fmt.Printf("Len -> %v, Cap -> %v", len(prices), cap(prices))

	// fmt.Println(prices)
	// fmt.Println(productNames)

	// fmt.Println(prices[3])
	// fmt.Println(productNames[0])

	//featuredPrices := prices[0:3]
	//featuredPrices := prices[:3] //0 - 2 index
	featuredPrices := prices[5:] // 200 - 600
	//featuredPrices[2] = 4000.4
	fmt.Println(featuredPrices)
	fmt.Println(len(featuredPrices),cap(featuredPrices))

	highlightedPrices := featuredPrices[:1] //200 - 600
	fmt.Println(highlightedPrices)

	fmt.Println(len(highlightedPrices),cap(highlightedPrices))
}