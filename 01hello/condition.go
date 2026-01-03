package main

import "fmt"

func conditionInGo() {
	var name string = "Muntasir"
	var age int = 25
	fmt.Println("Hello, " + name + "!")
	fmt.Printf("You are %d years old.\n", age)
	ageNext := 25
	fmt.Printf("Next year, you will be %d years old.\n", ageNext)
	fmt.Println("Goodbye, ", name, "!")

	if age < 18 {
		fmt.Println("You are still a minor.")
	} else {
		fmt.Println("You are an adult now.")
	}
	if age == 25 {
		fmt.Println("You are exactly 30 years old!")
	} else {
		fmt.Println("You are not 30 years old.")
	}
	if age >= 30 {
		fmt.Println("You are 30 or older.")
	} else {
		fmt.Println("You are younger than 30.")
	}
}
