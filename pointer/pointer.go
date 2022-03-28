package main

import "fmt"

func zero(value *int){
	*value=0
}

func main() {
	i := 1
	fmt.Println("value", i)
	zero(&i)
	fmt.Println("fromzero",i)
}