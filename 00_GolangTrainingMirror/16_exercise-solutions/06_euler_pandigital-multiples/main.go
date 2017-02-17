package main

import "fmt"

func main() {
	greatest := max(1, 2, 3, 4, 5, 6, 7, 8, 9)
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
