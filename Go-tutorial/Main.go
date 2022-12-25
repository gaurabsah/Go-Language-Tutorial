package main

import "fmt"

var name = "Gaurab"

const a int = 4
const b int = 5
const sum int = a + b

var product int = a * b

func main() {

	fmt.Println("Hello", name)

	fmt.Println("sum", sum)

	fmt.Println("Product", product)

	// incrementing by 1
	product++
	fmt.Println("New Product", product)
}
