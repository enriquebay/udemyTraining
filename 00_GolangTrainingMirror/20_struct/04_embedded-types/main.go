package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}

type doubleZero struct {
	person
	LicenseToKiil bool
}

func main() {
	p1 := doubleZero{
		person: person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		LicenseToKiil: true,
	}
	p2 := doubleZero{
		person: person{
			First: "Miss",
			Last:  "HoneyPenny",
			Age:   19,
		},
		LicenseToKiil: false,
	}

	fmt.Println(p1.First, p1.Last, p1.Age, p1.LicenseToKiil)
	fmt.Println(p2.First, p2.Last, p2.Age, p2.LicenseToKiil)
}
