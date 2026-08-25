package lists

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
	prices := []float64{10.99,8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	prices =append(prices, 5.99, 12.99, 29.99, 100.10)
	prices = prices[1:]
	fmt.Println(prices)

	discountPrices := []float64{101.99,80.99,20.59}
	prices = append(prices, discountPrices...)

	fmt.Println(prices)
}

// func main() {
// 	//var productNames [4]string = [4]string{"Spoon"}
// 	prices := [7]float64{100,200,300,400,500,600,700}
// 	fmt.Printf("Len -> %v, Cap -> %v", len(prices), cap(prices))

// 	fmt.Println(prices)
	// fmt.Println(productNames)

	// fmt.Println(prices[3])
	// fmt.Println(productNames[0])

	//featuredPrices := prices[1:4] //200-400
	//featuredPrices := prices[:3] //0 - 2 index
	// featuredPrices := prices[5:] // 200 - 600
	//featuredPrices[2] = 4000.4
// 	fmt.Println(featuredPrices)
// 	fmt.Println(len(featuredPrices),cap(featuredPrices))

// 	highlightedPrices := featuredPrices[:1] //200 - 600
// 	fmt.Println(highlightedPrices)

// 	fmt.Println(len(highlightedPrices),cap(highlightedPrices))
// }