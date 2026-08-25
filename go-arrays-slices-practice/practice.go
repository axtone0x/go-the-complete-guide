package main

import "fmt"

type Product struct {
	title string
	id string
	price float64
}

func main() {
	// 1)
	hobbies := [3]string{"Painting","Swimming","Dancing"}

	// 2)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	// 3)
	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)

	// 4)
	lastHobbies := mainHobbies[1:3]
	fmt.Println(lastHobbies)

	//5 
	courseGoals := []string{"Learn Go!","Learn all the basics"}
	fmt.Println(courseGoals)

	//6
	courseGoals[1] = "Learn all the details!"
	courseGoals = append(courseGoals, "Learn all the basics!")
	fmt.Println(courseGoals)

	products := []Product{{title: "Spoons", id: "AB0029X1", price: 157.80 },{title: "Forks", id: "AB0065X0", price: 177.18 }}
	fmt.Println(products)

	products = append(products, Product{title: "Umbrella", id:"AB0034X2", price: 9.99})
	fmt.Println(products)
}
