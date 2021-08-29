package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	alex.contact = contactInfo{email: "alexAnderson@gmail.com", zipCode: 94000}

	alex.updateName("Peter")

	alex.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
