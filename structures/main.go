package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo

	// This also works to create a field with name contactInfo and type contactInfo:
	// contactInfo
}

func main() {
	// Ways to instantiate a struct
	// matheus := person{"Matheus", "Araujo"}
	
	// matheus := person{firstName: "Matheus", lastName: "Araujo"}

	// var matheus person
	// matheus.firstName = "Matheus"
	// matheus.lastName = "Araujo"

	matheus := person{
		firstName: "Matheus",
		lastName: "Araujo",
		contact: contactInfo{
			email: "matheus@test.com",
			zipCode: 12345,
		},
	}

	matheusPointer := &matheus
	matheusPointer.updateName("Matias")
	matheus.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
