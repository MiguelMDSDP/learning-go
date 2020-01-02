package main

import "fmt"

// ShoppingList type simulates a shopping list
type ShoppingList []string

func main() {
	list := make(ShoppingList, 6)

	list[0] = "Lettuce"
	list[1] = "Cucumber"
	list[2] = "Olive Oil"
	list[3] = "Tuna Fish"
	list[4] = "Chicken"
	list[5] = "Chocolate"

	fmt.Println(list)

	vegetables, meat, others := list.Categorize()
	fmt.Println("Vegetables: ", vegetables)
	fmt.Println("Meat: ", meat)
	fmt.Println("Others: ", others)
}

// Categorize - Function that separates products into categories
func (list ShoppingList) Categorize() ([]string, []string, []string) {
	var vegetables, meat, others []string

	for _, e := range list {
		switch e {
		case "Lettuce", "Cucumber":
			vegetables = append(vegetables, e)
		case "Tuna Fish", "Chicken":
			meat = append(meat, e)
		case "Olive Oil", "Chocolate":
			others = append(others, e)
		}
	}

	return vegetables, meat, others
}
