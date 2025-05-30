package main

import "fmt"

func mapInGo() {
	fmt.Println("This is a map in Go:")
	ageMap := map[string]int{
		"Alice":   30,
		"Bob":     25,
		"Charlie": 35,
		"David":   28,
		"Eve":     22,
	}

	for name, age := range ageMap {
		fmt.Printf("%s is %d years old.\n", name, age)
	}
}
