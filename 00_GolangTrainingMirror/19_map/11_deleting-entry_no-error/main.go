package main

import "fmt"

func main() {

	myGreeting := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos días!",
		3: "Buongiorno!",
	}

	fmt.Println(myGreeting)
	delete(myGreeting, 7)
	fmt.Println(myGreeting)
}
