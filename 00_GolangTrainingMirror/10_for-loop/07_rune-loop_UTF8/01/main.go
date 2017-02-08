package main

import "fmt"

func main()  {
	for i := 250; i <= 300; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
	foo := "a"
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)
}