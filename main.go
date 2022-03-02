package main

import "fmt"

// create a person struct
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {

	// Different ways of implement a struct:
	// 1. Use the properties positions
	name1 := person{"Bill", "Gates", contactInfo{"", 12}}
	name1.print()

	// 2. Set the values fir specific properties
	name2 := person{lastName: "Jobs", firstName: "Steve"}
	name2.print()

	// 3. Using var
	var name3 person
	name3.firstName = "Elon"
	name3.lastName = "Musk"
	name3.print()
}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}
