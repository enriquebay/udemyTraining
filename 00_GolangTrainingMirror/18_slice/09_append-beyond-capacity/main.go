package main

import "fmt"

func main() {

	greeting := make([]string, 3, 5)

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "Huevos días!"

	greeting = append(greeting, "Supradahm!")
	greeting = append(greeting, "Zao an!")
	greeting = append(greeting, "Yo yo ma")
	greeting = append(greeting, "Yo ya mi")

	fmt.Println(greeting[6])
	fmt.Println(len(greeting))
	fmt.Println(cap(greeting))

}
