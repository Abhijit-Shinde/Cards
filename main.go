package main

import "fmt"

type Person struct {
	firstName   string
	lastName    string
	age         int
	skills      []string
	isMarried   bool
	contactInfo Contact
}

type Contact struct {
	email        string
	phoneNumber  int
	aadharNumber int
}

func main() {
	//isEvenOrOdd([]int {1,2,3,4,5,6,7,8,9,10})
	var person2 Person
	person2.age = 34
	person := Person{firstName: "Abhijit", lastName: "Shinde", age: 24, skills: []string{"abd", "124"}, isMarried: false,
		contactInfo: Contact{email: "abc@gmail.com", phoneNumber: 13525, aadharNumber: 226262}}

	person.updateFirstName("Abhszit")
	person.print()
}

func (p Person) print() {
	fmt.Printf("%+v", p)
}

func (p *Person) updateFirstName(newFirstName string) {
	p.firstName = newFirstName
}
