package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output () {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5) //initial length

	userNames[0] = "Julie"
	userNames[1] = "Robert"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	// courseRatings := map[string]float64{}
	// courseRatings["go"] = 5.0
	// courseRatings["react"] = 4.5

	// fmt.Println(courseRatings)

	courseRatings := make(floatMap,3)

	courseRatings["go"] = 5.0
	courseRatings["react"] = 4.5
	courseRatings["kafka"] = 4.8
	courseRatings["Docker"] = 4.0

	courseRatings.output()

	for index, value := range userNames {
		fmt.Println(index,value)
	}

	for key, value := range courseRatings {
		fmt.Println(key,value)
	}
}