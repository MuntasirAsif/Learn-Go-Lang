package main

import (
	"fmt"
	"sort"
)

func main(){
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("Original slice:", s)

	s = append(s, 6, 7, 8)
	fmt.Println("After appending elements:", s)

	s = append(s[:2], s[6:]...)
	fmt.Println("After removing elements at index 2 - 5:", s)

	highScore := make([]int , 5)

	highScore[0] = 100
	highScore[1] = 453
	highScore[2] = 22
	highScore[3] = 4003
	highScore[4] = 500

	fmt.Println("Initial high scores:", highScore,"First slices",s)
	highScore = append(highScore, 8000, 700)
	fmt.Println("High scores:", highScore)
	sort.Ints(highScore)
	fmt.Println("Sorted high scores:", highScore)
}