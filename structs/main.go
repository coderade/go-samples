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

	var name4 person
	name4.firstName = "Steve"
	name4.lastName = "Wozniak"
	//name4Pointer := &name4
	//fmt.Println("pointer: ", name4Pointer)
	name4.updateName("Billy") //pointer shortcut
	name4.print()
}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
