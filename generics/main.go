package main

import "fmt"

// func printSlice[T int | string](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

type Stack[T any] struct {
	elements []T
}

func main() {
	myStack := Stack[int]{
		elements: []int{1, 2, 3},
	}
	// nums := []int{1, 2, 3}

	// names := []string{"golang", "java", "c"}
	// printSlice(names)
	fmt.Println(myStack)
}
