package main

import "fmt"

func main() {

	greeting := []string{
		"Good morning!",
		"Bonjour!",
		"Huevos Días!",
		"Buongiorno!",
		"Ohayo!",
		"Selamat pagi!",
		"Gutten morgen!",
	}

	for i, currentEntrey := range greeting {
		fmt.Println(i, currentEntrey)
	}

	for j := 0; j < len(greeting); j++ {
		fmt.Println(greeting[j])

	}
}
