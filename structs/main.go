package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	//* alex := person{"Alex", "Anderson"}
	//* alex := person{firstName: "Alex", lastName: "Anderson"}
	//* var alex person
	//* alex.firstName = "Alex"
	//* alex.lastName = "Anderson"

	//* fmt.Println(alex)
	//* fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipcode: 57679,
		},
	}

	//* jimPointer := &jim
	//* jimPointer.updateFirstName("Jimmy")
	//* fmt.Printf("%p", jimPointer)
	jim.updateFirstName("Jimmy")

	jim.print()
}

func (person *person) updateFirstName(firstName string) {
	//* (*pointerToPerson).firstName = firstName
	person.firstName = firstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
