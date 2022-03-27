package main

import "fmt"

func plus(value1 int, value2 int) int {
	return value1 + value2
}

func main() {
	result := plus(12, 13)
	fmt.Println(result)
}