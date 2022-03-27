package main

import "fmt"

var product = make(map[string]float64)

func main() {
	fmt.Println("product =", product)

		//add
	product["Macbook"] = 40000
	product["Iphone"] = 30000
	product["Ipad"] = 20000
	fmt.Println("product",product)
		//del
	delete(product,"Ipad")
	fmt.Println("produc",product)

		//update
	product["Iphone"]=20000
	fmt.Println("product",product)

	value1 := product["Iphone"]
	fmt.Println("value1",value1)

	courseName := map[string]string{"101":"Java","102":"Python","103":"C#"}
	fmt.Println("CourseName",courseName)


}