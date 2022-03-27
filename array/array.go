package main

import "fmt"

var productName [4]string
var price [4]float32
func main() {
	productName[0]="Macbook"
	productName[1]="Ipad"
	price := [4]float32{60000,20000}
	fmt.Println(productName)
	fmt.Println(price)
}