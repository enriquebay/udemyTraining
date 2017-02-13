package main

import "fmt"

func main() {

	age := 44
	fmt.Println(age)
	changeMe(age)
}

func changeMe(z int) {
	fmt.Println(z)
	z = 24
}
