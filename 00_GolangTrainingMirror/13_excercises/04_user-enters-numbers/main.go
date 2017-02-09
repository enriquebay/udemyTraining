package main

import "fmt"

func main() {
	var large int
	var small int
	fmt.Println("Enter large number: ")
	fmt.Scan(&large)
	fmt.Println("Enter small number: ")
	fmt.Scan(&small)
	fmt.Println(large, "/", small, " = ", large / small)

}
