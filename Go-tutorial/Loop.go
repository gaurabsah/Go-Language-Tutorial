package main

import "fmt"

func main() {
	var i int
	var arr = [5]int{20, 10, 33, 44, 22}

	for i = 0; i < 5; i++ {
		fmt.Println(arr[i])
	}

	for true {
		fmt.Printf("This loop will run forever.\n")
	}
}
