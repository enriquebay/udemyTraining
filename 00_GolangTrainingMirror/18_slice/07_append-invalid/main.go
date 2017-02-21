package main

import "fmt"

func main() {

	greeting := make([]string, 3, 5)

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "Huevos días!"
	greeting[3] = "Supradahm!"

	fmt.Println(greeting[2])
}
