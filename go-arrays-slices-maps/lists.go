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
	var productNames [4]string = [4]string{"Spoon"}
	prices := [4]float64{10.9,9.99,15.25,0.77}

	fmt.Println(prices)
	fmt.Println(productNames)

	fmt.Println(prices[3])
	fmt.Println(productNames[3])
}