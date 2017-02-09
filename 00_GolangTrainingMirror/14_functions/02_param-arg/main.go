package main

import "fmt"

func main() {
	greet ("Joane")
	greet ("John")
}

func greet(name string)  {
	fmt.Println(name)
}