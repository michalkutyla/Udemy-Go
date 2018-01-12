package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo // or just contactInfo if we want the field to have same name as struct name
}

func main() {
	/*
		// Different ways of creating struct instances

		alex := person{"Alex", "Anderson"} // bad - it depends on the order of fields
		bill := person{firstName: "Bill", lastName: "Clinton"}
		var mike person
		mike.firstName = "Mike"
		mike.lastName = "Tyson"
		fmt.Println(alex, bill, mike) // prints {Alex Anderson} {Bill Clinton} {Mike Tyson}

		fmt.Printf("%+v", alex) //prints {firstName:Alex lastName:Anderson}%
	*/

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000, // <- this coma is necessary
		}, // <- this coma is necessary
	}

	// jimPointer := &jim
	// jimPointer.updateName("jimmy")
	jim.updateName("jimmy") // shortcut to above two lines
	jim.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
