package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	variableInGo()
	conditionInGo()

	// Using a while-like loop for the first loop
	i := 0
	for i < 5 {
		if i == 2 {
			fmt.Println("We are at the second iteration!")
			i++
			// Breaking out of the loop when i is 2
			continue
		}
		fmt.Println("This is iteration number", i)
		i++
	}

	// Using a while-like loop for the second loop
	k := 0
	for k < 5 {
		if k == 3 {
			// Breaking out of the loop when k is 3
			// This is the fourth iteration

			fmt.Println("We are at the fourth iteration!")
			//k++
			break
		}
		fmt.Println("This is iteration number", k)
		k++
	}

	printMessage()
	printInteger(42)

	var nums [3]int = [3]int{1, 2, 3}
	fmt.Println("Array elements:")
	for i, num := range nums {
		fmt.Printf("Element at index %d: %d\n", i, num)
	}

	mapInGo()
}

func printMessage() string {
	fmt.Println("This is a message from the printMessage function.")
	return "Message printed successfully."
}

func printInteger(i int) int {
	fmt.Println("The integer is:", i)
	return i
}
