package main

import "fmt"

func changeNum(num *int) {
	*num = 5
}

func main() {
	num := 1
	changeNum(&num)
	fmt.Println(num)
}
