package main

import "fmt"

func main() {
	greatest := max(5, 8, 12, 20)
	fmt.Println(greatest)
}

func max(numbers ...int) int {
	var largest int
	for _, v := range numbers {
		if v > largest {
			largest = v
		}
	}
	return largest
}
