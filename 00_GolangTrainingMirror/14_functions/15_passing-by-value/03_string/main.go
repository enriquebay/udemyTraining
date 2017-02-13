package main

import "fmt"

func main() {

	name := "Enrique"
	fmt.Println(name)

	changeMe(name)

	fmt.Println(name)

}

func changeMe(z string) {
	fmt.Println(z)
	z = "Rocky"
	fmt.Println(z)
}
