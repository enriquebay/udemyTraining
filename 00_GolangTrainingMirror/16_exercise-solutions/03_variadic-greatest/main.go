package main

import "fmt"

func main() {
	greatest := max(1, 11, 111, 1111)
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
