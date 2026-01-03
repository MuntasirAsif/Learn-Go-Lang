package main

import "fmt"

func main() {
	var fruits [3]string

	fruits[0] = "Apple"
	fruits[1] = "Mango"
	fruits[2] = "Banana"
	fmt.Println("Fruits list:", fruits)

	var vegitables = [5]string{"Potato", "Tomato", "Cabbage"}
	fmt.Println("Vegitables list length:", len(vegitables))
}